package bankbuster

import (
	"bytes"
	"io"
	"net/http"

	json "github.com/bytedance/sonic"
)

var (
	DefaultHTTPMethod string      = http.MethodPost
	DefaultHTTPHeader http.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	// Sanity check if transport implements http.RoundTripper interface
	_ http.RoundTripper = &Transport{}
)

type Transport struct {
	proxyURL string
	client   *http.Client
}

func NewTransport(proxyURL string) *Transport {
	return &Transport{
		proxyURL: proxyURL,
		client:   &http.Client{},
	}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	requestSerialized, err := SerializeRequest(req)
	if err != nil {
		return nil, err
	}

	requestPayload, err := json.Marshal(requestSerialized)
	if err != nil {
		return nil, err
	}

	requestProxied, err := http.NewRequestWithContext(req.Context(), DefaultHTTPMethod, t.proxyURL, bytes.NewReader(requestPayload))
	if err != nil {
		return nil, err
	}
	requestProxied.Header = DefaultHTTPHeader

	responseProxied, err := t.client.Do(requestProxied)
	if err != nil {
		return nil, err
	}
	defer responseProxied.Body.Close()

	if responseProxied.StatusCode != http.StatusOK {
		return nil, nil
	}

	buf, err := io.ReadAll(responseProxied.Body)
	if err != nil {
		return nil, err
	}

	var res HTTPResponse
	if err = json.Unmarshal(buf, &res); err != nil {
		return nil, err
	}

	return DeserializeRequest(&res)
}
