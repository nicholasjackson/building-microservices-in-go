package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(nil)

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{})

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func setupTest(data interface{}) (*http.Request, *httptest.ResponseRecorder, SearchHandler) {
	h := SearchHandler{}
	rw := httptest.NewRecorder()

	if data == nil {
		return httptest.NewRequest("POST", "/search", nil), rw, h
	}

	body, _ := json.Marshal(searchRequest{})
	return httptest.NewRequest("POST", "/search", bytes.NewReader(body)), rw, h
}
