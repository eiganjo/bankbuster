package bankbuster

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type HTTPResponse struct {
	Status     string            `json:"status"`
	StatusCode int               `json:"statusCode"`
	Body       string            `json:"body,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
}

func DeserializeRequest(res *HTTPResponse) (*http.Response, error) {
	if res == nil {
		return nil, fmt.Errorf("can not deserialize, response is nil")
	}

	var (
		bodyBytes []byte
		err       error
	)
	if res.Body != "" {
		bodyBytes, err = base64.StdEncoding.DecodeString(res.Body)
		if err != nil {
			return nil, err
		}
	}

	resp := &http.Response{
		StatusCode: res.StatusCode,
		Header:     http.Header{},
		Body:       io.NopCloser(bytes.NewBuffer(bodyBytes)),
	}

	for k, v := range res.Headers {
		resp.Header.Add(k, v)
	}

	return resp, nil
}
