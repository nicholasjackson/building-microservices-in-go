package main

import (
	"fmt"
	"log"
	"net/rpc"

	entities "github.com/nicholasjackson/building-microservices-in-go/chapter2/rpc"
)

const port = 1234

func main() {
	client := CreateClient()
	defer client.Close()

	reply := PerformRequest(client)
	fmt.Println(reply.Message)
}

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) entities.HelloWorldResponse {
	args := &entities.HelloWorldRequest{Name: "World"}
	var reply entities.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
