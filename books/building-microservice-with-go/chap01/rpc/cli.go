package main

import (
	"fmt"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc/client"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
