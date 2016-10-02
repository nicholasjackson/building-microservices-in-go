package main

import (
	"fmt"

	"github.com/nicholasjackson/building-microservices-in-go/chapter1/rpc_http/client"
	"github.com/nicholasjackson/building-microservices-in-go/chapter1/rpc_http/server"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
