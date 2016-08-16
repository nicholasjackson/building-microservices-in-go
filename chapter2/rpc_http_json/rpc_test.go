package main

import (
	"bytes"
	"net/http"
	"testing"
)

func BenchmarkHelloWorldHandlerJSONRPC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r, _ := http.Post(
			"http://localhost:1234",
			"application/json",
			bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
		)
		r.Body.Close()
		/*
			data, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()

			fmt.Println(string(data)) // {"id":1,"result":{"message":"Hello World"},"error":null}
		*/
	}
}

func init() {
	// start the server
	go server()
}
