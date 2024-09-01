package main

import (
    "io"
    "net/http"

    json "github.com/bytedance/sonic"
    "github.com/GoogleCloudPlatform/functions-framework-go/functions"

    "github.com/eiganjo/bankbuster"
)

func init() {
    functions.HTTP("bankbusterHTTPProxy", HTTPProxyHandler)
}

func HTTPProxyHandler(w http.ResponseWriter, r *http.Request) {
    	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var request bankbuster.HTTPRequest

	if err = json.Unmarshal(body, &request); err != nil {
	    return
	}





}
