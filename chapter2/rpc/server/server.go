package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	entities "github.com/nicholasjackson/building-microservices-in-go/chapter2/rpc"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *entities.HelloWorldRequest, reply *entities.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	helloWorld := new(HelloWorldHandler)
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
