package main

import (
	"fmt"

	"github.com/nicholasjackson/building-microservices-in-go/chapter2/rpc/client"
	"github.com/nicholasjackson/building-microservices-in-go/chapter2/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
