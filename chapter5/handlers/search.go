package handlers

import (
	"encoding/json"
	"net/http"
)

type searchRequest struct {
	// Query is the text search query that will be executed by the handler
	Query string `json:"query"`
}

// SearchHandler is an http handler for our microservice
type SearchHandler struct {
}

func (s *SearchHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	request := new(searchRequest)
	err := decoder.Decode(request)
	if err != nil || len(request.Query) < 1 {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

}
