package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/alexcesaro/statsd"
	"github.com/nicholasjackson/building-microservices-in-go/chapter7/server/entities"
	"github.com/nicholasjackson/building-microservices-in-go/chapter7/server/httputil"
)

type helloWorldHandler struct {
	statsd *statsd.Client
	logger *logrus.Logger
}

// NewHelloWorldHandler creates a new handler with the given logger and
// statsD client
func NewHelloWorldHandler(statsd *statsd.Client, logger *logrus.Logger) http.Handler {
	return &helloWorldHandler{statsd: statsd, logger: logger}
}

func (h *helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	timing := h.statsd.NewTiming()

	name := r.Context().Value("name").(string)
	response := entities.HelloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)

	h.statsd.Increment(helloworldSuccess)

	message := httputil.SerialzableRequest{r}
	h.logger.WithFields(logrus.Fields{
		"handler": "HelloWorld",
		"status":  http.StatusOK,
		"method":  r.Method,
	}).Info(message.ToJSON())

	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	timing.Send(helloworldTiming)
}
