package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nicholasjackson/building-microservices-in-go/chapter4/benchmark/data"
)

type searchRequest struct {
	// Query is the text search query that will be executed by the handler
	Query string `json:"query"`
}

type searchResponse struct {
	Kittens []data.Kitten `json:"kittens"`
}

// Search is an http handler for our microservice
type Search struct {
	DataStore data.Store
}

func (s *Search) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	request := new(searchRequest)

	err := decoder.Decode(request)
	if err != nil || len(request.Query) < 1 {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	kittens := s.DataStore.Search(request.Query)

	encoder := json.NewEncoder(rw)
	encoder.Encode(searchResponse{Kittens: kittens})
}
