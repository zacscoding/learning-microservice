package main

import (
	"fmt"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc_http/client"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc_http/server"
)

func main() {
	server.StartServer() // TODO : 블로킹 됨

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println("rpc result : ", reply.Message)
}
