package main

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func BenchmarkDialHTTP(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
		if err != nil {
			log.Fatal("dialing:", err)
		}
		client.Close()
	}
}

func BenchmarkHelloWorldHandlerHTTP(b *testing.B) {
	b.ResetTimer()

	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	for i := 0; i < b.N; i++ {
		args := &HelloWorldRequest{Name: "World"}

		var reply HelloWorldResponse
		err = client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	}

	client.Close()

}

func init() {
	// start the server
	go server()
}
