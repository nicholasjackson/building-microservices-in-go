package main

import (
	"fmt"
	"log"
	"net/http"
)

struct HTTPResponse {
	Message string `"json:message"`
}

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Printf("Server running on port %i\n", 8080)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
