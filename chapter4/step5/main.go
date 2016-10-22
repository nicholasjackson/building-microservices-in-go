package main

import (
	"log"
	"net/http"

	"github.com/nicholasjackson/building-microservices-in-go/chapter4/step5/data"
	"github.com/nicholasjackson/building-microservices-in-go/chapter4/step5/handlers"
)

func main() {
	store, err := data.NewMongoStore("localhost")
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Search{DataStore: store}
	err = http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
