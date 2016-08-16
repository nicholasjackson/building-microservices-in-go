package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

const port = 1234

func main() {
	if os.Args[1] == "server" {
		server()
	} else {
		client()
	}
}

func server() {
	helloWorld := new(HelloWorldHandler)
	rpc.Register(helloWorld)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, nil)
}

func client() {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	defer client.Close()

	args := &HelloWorldRequest{Name: "World"}

	var reply HelloWorldResponse
	err = client.Call("HelloWorldHandler.HelloWorld", args, &reply)

	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Println(reply.Message)
}

type HelloWorldRequest struct {
	Name string
}

type HelloWorldResponse struct {
	Message string
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *HelloWorldRequest, reply *HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
