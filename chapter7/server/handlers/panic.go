package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Sirupsen/logrus"
	"github.com/alexcesaro/statsd"
)

type panicHandler struct {
	logger *logrus.Logger
	statsd *statsd.Client
	next   http.Handler
}

func (p *panicHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			p.logger.WithFields(
				logrus.Fields{
					"handler": "panic",
					"status":  http.StatusInternalServerError,
					"method":  r.Method,
					"path":    r.URL.Path,
					"query":   r.URL.RawQuery,
				},
			).Error(fmt.Sprintf("Error: %v\n%s", err, debug.Stack()))

			rw.WriteHeader(http.StatusInternalServerError)
		}
	}()

	p.next.ServeHTTP(rw, r)
}

func NewPanicHandler(statsd *statsd.Client, logger *logrus.Logger, next http.Handler) http.Handler {
	return &panicHandler{statsd: statsd, logger: logger, next: next}
}
