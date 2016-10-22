package main

import (
	"log"
	"net/http"

	"github.com/nicholasjackson/building-microservices-in-go/chapter4/step4/data"
	"github.com/nicholasjackson/building-microservices-in-go/chapter4/step4/handlers"
)

func main() {
	handler := handlers.Search{DataStore: &data.MemoryStore{}}
	err := http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
