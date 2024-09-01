package bankbuster

import (
	"net/http"
	"sync"
)

var _ http.RoundTripper = &RoundRobinTripperTransport{}

type RoundRobinTripperTransport struct {
	transports      []http.RoundTripper
	transportsMutex sync.Mutex
	transportsIndex int
}

func NewRoundRobinTripperTransport(transports []http.RoundTripper) *RoundRobinTripperTransport {
	return &RoundRobinTripperTransport{transports: transports}
}

func (r *RoundRobinTripperTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r.transportsMutex.Lock()
	transport := r.transports[r.transportsIndex]
	r.transportsIndex = (r.transportsIndex + 1) % len(r.transports)
	r.transportsMutex.Unlock()

	return transport.RoundTrip(req)
}
