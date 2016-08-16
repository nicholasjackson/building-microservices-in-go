package main

import (
	"fmt"

	"github.com/nicholasjackson/building-microservices-in-go/chapter2/rpc_http/client"
	"github.com/nicholasjackson/building-microservices-in-go/chapter2/rpc_http/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
