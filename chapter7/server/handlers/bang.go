package handlers

import "net/http"

type bangHandler struct {
}

func (b *bangHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	panic("Somethings gone wrong again")
}

// NewBangHandler creates a handler which panics
func NewBangHandler() http.Handler {
	return &bangHandler{}
}
