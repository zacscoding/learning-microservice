package server

import (
	"fmt"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc_http/contract"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const Port = 12345

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	fmt.Println("HelloWorldHandler.HelloWorld is called")
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer() {
	helloWorld := new(HelloWorldHandler)

	rpc.Register(helloWorld)
	// HTTP 를 통해 RPC를 사용하기에 반드시 필요
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port : %s", err))
	}

	log.Printf("Server starting on port %v\n", Port)

	http.Serve(l, nil)
}
