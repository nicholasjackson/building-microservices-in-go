package main

import "github.com/nicholasjackson/building-microservices-in-go/chapter2/rpc/server"

//curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:1234
func main() {
	server.StartServer()
}
