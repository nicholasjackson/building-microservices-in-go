package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/context"
)

type key int

type requestContext struct {
	context *context.Context
	cancel  context.CancelFunc
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

const nameKey key = 0

var contexts = map[*http.Request]requestContext{}
var contextLock sync.Mutex
var ctx = context.Background()

// returns the Context for our request, creates a new Context if one does not exist
func contextForRequest(r *http.Request) *context.Context {
	contextLock.Lock()
	defer contextLock.Unlock()

	c, ok := contexts[r]

	if !ok {
		c, cancel := context.WithCancel(ctx) // copy root and add to map
		contexts[r] = requestContext{context: &c, cancel: cancel}
		return &c
	}

	return c.context // context has already been created
}

// deletes the context for our request, must be called or we will leak memory
func deleteContextForRequest(r *http.Request) {
	contextLock.Lock()
	defer contextLock.Unlock()

	contexts[r].cancel()

	delete(contexts, r)
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	} else {
		ctx := contextForRequest(r)
		*ctx = context.WithValue(*ctx, nameKey, request.Name)

		defer deleteContextForRequest(r) // cleanup
		h.next.ServeHTTP(rw, r)
	}
}

type helloWorldHandler struct {
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := *contextForRequest(r)
	name := ctx.Value(nameKey).(string)

	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
