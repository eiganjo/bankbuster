package bankbuster

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
)

type HTTPRequest struct {
	URL     string            `json:"url"`
	Method  string            `json:"method,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}

func SerializeRequest(req *http.Request) (*HTTPRequest, error) {
	// Allocate on demand
	var headers map[string]string
	if len(req.Header) > 0 {
		headers = make(map[string]string, len(req.Header))
		for k, v := range req.Header {
			if len(v) > 0 {
				headers[k] = v[0]
			}
		}
	}

	// Read the body on demand
	var body string
	if req.Body != nil {
		// Use global sync.Pool for buffer
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		defer bufferPool.Put(buf)

		if _, err := buf.ReadFrom(req.Body); err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(buf)
		body = base64.StdEncoding.EncodeToString(buf.Bytes())
	}

	return &HTTPRequest{
		Method:  req.Method,
		URL:     req.URL.String(),
		Headers: headers,
		Body:    body,
	}, nil
}
