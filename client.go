package bankbuster

import (
	"net/http"
)

func NewHTTPClientRoundRobin(bankbusterProxyURLs ...string) *http.Client {
	var client *http.Client

	switch len(bankbusterProxyURLs) {
	case 0:
		client = &http.Client{}
	case 1:
		client = &http.Client{
			Transport: NewTransport(bankbusterProxyURLs[0]),
		}
	default:
		var transports []http.RoundTripper

		for _, url := range bankbusterProxyURLs {
			transports = append(transports, NewTransport(url))
		}

		client = &http.Client{
			Transport: NewRoundRobinTripperTransport(transports),
		}
	}

	return client
}
