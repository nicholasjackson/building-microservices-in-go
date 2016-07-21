package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloWorldResponse struct {
	Message string
}

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Printf("Server running on port %s\n", 8080)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
