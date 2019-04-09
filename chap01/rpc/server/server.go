package server

import (
	"fmt"
	"github.com/zacscoding/learning-microservice-with-go/chap01/rpc/contract"
	"log"
	"net"
	"net/rpc"
)

const Port = 1234

func main() {
	log.Printf("Server starting on port %v\n", Port)
	StartServer()
}

func StartServer() {
	helloWorld := new(HelloWorldHandler)
	rpc.Register(helloWorld)

	// Listener 인터페이스를 구현하는 인스턴스를 리턴
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port : %s", err))
	}

	defer l.Close()

	for {
		conn, _ := l.Accept()
		log.Println("Accept new connection ", l.Addr())
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct {
}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	log.Println("HelloWorldHandler.HelloWorld is called... args : ", args)
	reply.Message = "Hello " + args.Name
	return nil
}
