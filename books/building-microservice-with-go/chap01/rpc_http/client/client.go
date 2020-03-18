package client

import (
	"fmt"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc_http/contract"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc_http/server"
	"log"
	"net/rpc"
)

func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", server.Port))

	if err != nil {
		log.Fatal("dialing : ", err)
	}

	fmt.Println("success to create http client ", client)

	return client
}

func PerformRequest(c *rpc.Client) contract.HelloWorldResponse {
	fmt.Println("PerformRequest is called")
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse
	err := c.Call("HelloWorldHandler.HelloWorld", args, &reply)

	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
