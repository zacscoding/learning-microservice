package client

import (
	"bytes"
	"encoding/json"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc_http_json/contract"
	"net/http"
)

func PerformRequest() contract.HelloWorldResponse {
	r, _ := http.Post(
		"http://localhost:12345",
		"application/json",
		bytes.NewBuffer([]byte(`{"id" : 1, "method" : "HelloWorldHandler.HelloWorld", "params" : [{"name", "world"}]`)),
	)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var response contract.HelloWorldResponse
	decoder.Decode(&response)

	return response
}
