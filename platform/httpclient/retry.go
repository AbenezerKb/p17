package httpclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func (s *httpclient) shouldRetry(err error, req *http.Request, resp *http.Response, numRetries int) (bool, string) {
	if numRetries >= int(s.retryCount) {
		return false, "max retries exceeded"
	}

	// Don't retry if the context was was canceled or its deadline was
	// exceeded.
	if req.Context() != nil && req.Context().Err() != nil {
		switch req.Context().Err() {
		case context.Canceled:
			return false, "context canceled"
		case context.DeadlineExceeded:
			return false, "context deadline exceeded"
		default:
			return false, fmt.Sprintf("unknown context error: %v", req.Context().Err())
		}
	}

	if err != nil {
		return true, ""
	}

	// 409 Conflict
	if resp.StatusCode == http.StatusConflict {
		return true, ""
	}

	// 429 Too Many Requests
	//
	// There are a few different problems that can lead to a 429. The most
	// common is rate limiting, on which we *don't* want to retry because
	// that'd likely contribute to more contention problems. However, some 429s
	// are lock timeouts, which is when a request conflicted with another
	// request or an internal process on some particular object. These 429s are
	// safe to retry.

	// if resp.StatusCode == http.StatusTooManyRequests {

	// }

	// 500 Internal Server Error
	//
	// We only bother retrying these for non-POST requests. POSTs
	// could cause a state change
	if resp.StatusCode >= http.StatusInternalServerError && req.Method != http.MethodPost {
		return false, ""
	}

	// 503 Service Unavailable
	if resp.StatusCode == http.StatusServiceUnavailable {
		return true, ""
	}

	return false, "response not known to be safe for retry"
}

func resetBodyReader(body *bytes.Buffer, req *http.Request) {
	if body != nil {
		reader := bytes.NewReader(body.Bytes())

		req.Body = nopReadCloser{reader}
		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body.Bytes())
			return nopReadCloser{reader}, nil
		}
	}
}
