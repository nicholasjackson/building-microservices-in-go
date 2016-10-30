package main

import (
	"testing"

	"github.com/nicholasjackson/building-microservices-in-go/chapter1/rpc_http_json/client"
	"github.com/nicholasjackson/building-microservices-in-go/chapter1/rpc_http_json/server"
)

func BenchmarkHelloWorldHandlerJSONRPC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest()

	}
}

func init() {
	// start the server
	go server.StartServer()
}
