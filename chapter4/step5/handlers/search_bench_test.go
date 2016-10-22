package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/nicholasjackson/building-microservices-in-go/chapter4/step5/data"
)

func BenchmarkSearchHandler(b *testing.B) {
	mockStore = &data.MockStore{}
	mockStore.On("Search", "Fat Freddy's Cat").Return([]data.Kitten{
		data.Kitten{
			Name: "Fat Freddy's Cat",
		},
	})

	search := Search{DataStore: mockStore}

	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("POST", "/search", bytes.NewReader([]byte(`{"query":"Fat Freddy's Cat"}`)))
		rr := httptest.NewRecorder()
		search.ServeHTTP(rr, r)
	}
}
