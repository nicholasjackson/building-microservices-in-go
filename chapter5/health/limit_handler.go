package health

import "net/http"

// LimitHandler is middleware which limits the current number of active
// connections that this handler can sustain.
// Once the current connections equal the max http.StatusTooManyRequests is
// returned
type LimitHandler struct {
	connections chan struct{}
	handler     http.Handler
}

// NewLimitHandler creates a new instance of the LimitHandler for the
// given parameters.
func NewLimitHandler(connections int, next http.Handler) *LimitHandler {
	cons := make(chan struct{}, connections)
	for i := 0; i < connections; i++ {
		cons <- struct{}{}
	}

	return &LimitHandler{
		connections: cons,
		handler:     next,
	}
}

func (l *LimitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	select {
	case <-l.connections:
		l.handler.ServeHTTP(rw, r)
		l.connections <- struct{}{} // release the lock
	default:
		http.Error(rw, "Busy", http.StatusTooManyRequests)
	}
}
