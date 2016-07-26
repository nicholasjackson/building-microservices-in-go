package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HelloWorldResponse struct {
	Message string `json:"message"`
}

type HelloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Starting server on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	var request HelloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	response := HelloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
