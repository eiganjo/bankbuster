package bankbuster

import (
	"context"
	"net/http"
)

type ProxyHandler func(ctx context.Context, req *HTTPRequest) (*HTTPResponse, error)

func NewProxyHandler(clients ...*http.Client) ProxyHandler {

    return nil
}
