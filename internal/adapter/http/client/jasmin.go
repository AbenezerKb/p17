package client

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"os"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
)

type jasmineClient struct {
	utils const_init.Utils
}

const (
	SendSMS      = `http://127.0.0.1:8080/secure/send`
	SendBatchSMS = `http://127.0.0.1:8080/secure/sendbatch`
)

type Response struct {
	BatchId      string `json:"batchId"`
	MessageCount int    `json:"messageCount"`
}

type BatchResponse struct {
	//BatchConfig string   `json:"batch_config""`
	Response Response `json:"data"`
}

type JasminClient interface {
	OutGoingSMS(ctx context.Context, sms *model.OutGoingSMS) (*string, error)
	BatchOutGoingSMS(ctx context.Context, sms *model.SMS) (*Response, error)
	IncomingSms()
}

// ClientInit initializes jasmine interfaces
func ClientInit(utils const_init.Utils) JasminClient {
	return jasmineClient{
		utils: utils,
	}

}

// OutGoingSMS send single SMS
func (j jasmineClient) OutGoingSMS(ctx context.Context, sms *model.OutGoingSMS) (*string, error) {
	client := &http.Client{}
	basic_auth := os.Getenv("BASIC_AUTH")
	basic := `Basic %s`
	payload := `{
		"to": "%s",
		"content": "%s"
	}`

	body := bytes.NewBuffer([]byte(fmt.Sprintf(payload, sms.To, sms.Content)))

	req, err := http.NewRequest("GET", SendSMS, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf(basic, basic_auth))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, nil
	}

	SMSResponse := ""
	err = json.Unmarshal(bdy, SMSResponse)
	if err != nil {

		return nil, err
	}

	return &SMSResponse, nil
}

func (j jasmineClient) IncomingSms() {

}

// BatchOutGoingSMS Sends SMS in Batch
func (j jasmineClient) BatchOutGoingSMS(ctx context.Context, sms *model.SMS) (*Response, error) {
	client := &http.Client{}
	basic_auth := os.Getenv("BASIC_AUTH")
	basic := `Basic %s`
	payload := `{
  "batch_config": {
    "callback_url": "http://127.0.0.1:7877/successful_batch",
    "errback_url": "http://127.0.0.1:7877/errored_batch"
      },
  "messages": [
    {
      "to": %s,
      "content": "%s"
    }
    
  ]
}`
	str := ""
	for _, v := range sms.To {
		str += `,"` + v + `"`
	}
	str = str[1:]
	bd := "[" + str + "]"
	body := bytes.NewBuffer([]byte(fmt.Sprintf(payload, bd, sms.Content)))

	req, err := http.NewRequest("POST", SendBatchSMS, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf(basic, basic_auth))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	bdy, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return nil, nil
	}

	SMSResponse := &BatchResponse{}
	err = json.Unmarshal(bdy, SMSResponse)
	if err != nil {

		return nil, err
	}

	return &SMSResponse.Response, nil
}
