package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"sync"
)

type key int

const nameKey key = 0

type requestContext struct {
	context *context.Context
	cancel  context.CancelFunc
}

var contexts = map[*http.Request]requestContext{}
var contextLock sync.Mutex
var ctx = context.Background()

// returns the Context for our request, creates a new Context if one does not exist
func contextForRequest(r *http.Request) *context.Context {
	contextLock.Lock()
	defer contextLock.Unlock()
	if c, ok := contexts[r]; ok {
		return c.context // context has already been created
	} else {
		c, cancel := context.WithCancel(ctx) // copy root and add to map
		contexts[r] = requestContext{context: &c, cancel: cancel}
		return &c
	}
}

// deletes the context for our request, must be called or we will leak memory
func deleteContextForRequest(r *http.Request) {
	contextLock.Lock()
	defer contextLock.Unlock()

	contexts[r].cancel()

	delete(contexts, r)
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type validationHandler struct {
	next http.Handler
}

type helloWorldHandler struct {
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func main() {
	handler := newValidationHandler(newHelloWorldHandler())

	http.Handle("/helloworld", handler)

	log.Printf("Starting server on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
	} else {
		ctx := contextForRequest(r)
		*ctx = context.WithValue(*ctx, nameKey, request.Name)

		h.next.ServeHTTP(rw, r)
	}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := *contextForRequest(r)
	defer deleteContextForRequest(r) // cleanup

	name := ctx.Value(nameKey).(string)

	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
