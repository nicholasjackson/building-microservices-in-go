package main

import "github.com/micro/go-micro/server"

func main() {
	server.Init(
		server.Name("bmigo.micro.kittens"),
		server.Version("1.0.0"),
	)


}