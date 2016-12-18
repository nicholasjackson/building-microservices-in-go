package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/alexcesaro/statsd"
	logstash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/nicholasjackson/building-microservices-in-go/chapter7/server/handlers"
)

const port = 8091

func main() {
	statsd, err := createStatsDClient(os.Getenv("STATSD"))
	if err != nil {
		log.Fatal("Unable to create statsD client")
	}

	logger, err := createLogger(os.Getenv("LOGSTASH"))
	if err != nil {
		log.Fatal("Unable to create logstash client")
	}

	setupHandlers(statsd, logger)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func setupHandlers(statsd *statsd.Client, logger *logrus.Logger) {
	handler := handlers.NewValidationHandler(
		statsd,
		logger,
		handlers.NewHelloWorldHandler(statsd, logger),
	)

	bangHandler := handlers.NewPanicHandler(
		statsd,
		logger,
		handlers.NewBangHandler(),
	)

	http.Handle("/helloworld", handler)
	http.Handle("/bang", bangHandler)
}

func createStatsDClient(address string) (*statsd.Client, error) {
	return statsd.New(statsd.Address(address))
}

func createLogger(address string) (*logrus.Logger, error) {
	retryCount := 0

	l := logrus.New()
	hostname, _ := os.Hostname()
	var err error

	// Retry connection to logstash incase the server has not yet come up
	for ; retryCount < 10; retryCount++ {
		hook, err := logstash.NewHookWithFields(
			"tcp",
			address,
			"kittenserver",
			logrus.Fields{"hostname": hostname},
		)

		if err == nil {
			l.Hooks.Add(hook)
			return l, err
		}

		log.Println("Unable to connect to logstash, retrying")
		time.Sleep(1 * time.Second)
	}

	return nil, err
}
