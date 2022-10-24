package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient interface {
	HttpClient() *http.Client
	DoRaw(headers map[string]string, body *bytes.Buffer, req *http.Request, responseObj, errObj interface{}) error
	NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error)
	Do(headers map[string]string, body interface{}, req *http.Request, responseObj, errObj interface{}) error
}

type httpclient struct {
	retryCount int
}

func InitHttpClient() HttpClient {
	return &httpclient{
		retryCount: 5,
	}
}

func (httpclient *httpclient) HttpClient() *http.Client {
	// Customize the Transport to have larger connection pool
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}

	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	httpClient := &http.Client{Transport: &defaultTransport}

	return httpClient

}

func (httpclient *httpclient) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	return req.WithContext(ctx), nil
}

func (HttpClient *httpclient) Do(headers map[string]string, body interface{}, req *http.Request, responseObj, errObj interface{}) error {
	if body == nil {
		return HttpClient.DoRaw(headers, nil, req, responseObj, errObj)
	}

	bodyString, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bodyBuffer := bytes.NewBufferString(string(bodyString))
	return HttpClient.DoRaw(headers, bodyBuffer, req, responseObj, errObj)
}

// with back off and reties
func (httpclient *httpclient) DoRaw(headers map[string]string, body *bytes.Buffer, req *http.Request, responseObj, errObj interface{}) error {
	var (
		retries     = 0
		backOffTime = 3 // seconds
	)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if headers["Content-Type"] == "" {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}

	if headers["Accept"] == "" {
		req.Header.Set("Accept", "application/json")
	}

	var resp *http.Response
	var err error

	// retry till success or max number of retries reached

	for {
		if retries != 0 {
			backOffTime *= retries
			time.Sleep(time.Duration(backOffTime) * time.Second)
		}

		resetBodyReader(body, req)
		resp, err = httpclient.HttpClient().Do(req)

		// if not successful
		if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
			// if request failed decrease retry count
			shouldRetry, noRetryReason := httpclient.shouldRetry(err, req, resp, retries)

			if !shouldRetry {
				fmt.Printf("Not retrying request: %v %v\n", noRetryReason, err)
				if err == nil {
					fmt.Println("Status-Code", resp.StatusCode)
				}
				break
			}

			retries += 1
		} else {
			break
		}
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Checking the response for status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// attach the error body of the error response from server
		_ = json.Unmarshal(data, errObj)

		return fmt.Errorf("response status indicates error - %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, responseObj)
}
