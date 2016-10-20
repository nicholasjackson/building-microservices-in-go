package main

import (
	"log"
	"net/http"

	"github.com/nicholasjackson/building-microservices-in-go/chapter5/handlers"
)

func main() {
	err := http.ListenAndServe(":2323", &handlers.SearchHandler{})
	if err != nil {
		log.Fatal(err)
	}
}
