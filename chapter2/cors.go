package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
