package client

import (
	"fmt"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc/contract"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc/server"
	"log"
	"net/rpc"
)

func CreateClient() *rpc.Client {
	// Client 생성
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", server.Port))
	if err != nil {
		log.Fatal("error : ", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	// 서버의 이름 붙여진 (named) 함수를 호출
	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error : ", err)
	}

	return reply
}
