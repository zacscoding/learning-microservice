package main

import (
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc_http_json/server"
)

func main() {
	server.StartServer()
	// curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:12345
}
