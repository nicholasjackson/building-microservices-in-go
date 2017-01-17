package validation

import (
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
)

// Request defines the input structure received by a http handler
type Request struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	URL   string `json:"url" validate:"url"`
}

var validate = validator.New()

func Handler(rw http.ResponseWriter, r *http.Request) {
	request := Request{}

	err := json.NewEncoder(rw).Encode(&request)
	if err != nil {
		http.Error(rw, "Invalid request object", http.StatusBadRequest)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
