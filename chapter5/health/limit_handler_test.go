package health

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTestHandler(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		<-r.Context().Done()
	})
}

func setup(ctx context.Context) (*httptest.ResponseRecorder, *http.Request) {
	r := httptest.NewRequest("GET", "/health", nil)
	r = r.WithContext(ctx)
	return httptest.NewRecorder(), r
}

func testCallsNextWhenConnectionsOK(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	handler := NewLimitHandler(2, newTestHandler(ctx))
	rw, r := setup(ctx)

	go handler.ServeHTTP(rw, r)
	cancel()
	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, http.StatusOK, rw.Code)
}

func TestReturnsBusyWhen0Connections(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	handler := NewLimitHandler(0, newTestHandler(ctx))
	rw, r := setup(ctx)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
	})
	handler.ServeHTTP(rw, r)

	assert.Equal(t, http.StatusTooManyRequests, rw.Code)
}

func TestReturnsOKWith2ConnnectionsAndConnectionLimit2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	handler := NewLimitHandler(2, newTestHandler(ctx))
	rw, r := setup(ctx)
	rw2, r2 := setup(ctx2)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
		cancel2()
	})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		handler.ServeHTTP(rw, r)
		waitGroup.Done()
	}()

	go func() {
		handler.ServeHTTP(rw2, r2)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if rw.Code != http.StatusOK && rw2.Code != http.StatusOK {
		t.Fatalf("Both requests should be OK, request 1: %v, request 2: %v", rw.Code, rw2.Code)
	}
}

func TestReturnsBusyWhenConnectionsExhausted(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	handler := NewLimitHandler(1, newTestHandler(ctx))
	rw, r := setup(ctx)
	rw2, r2 := setup(ctx2)

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
		cancel2()
	})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		handler.ServeHTTP(rw, r)
		waitGroup.Done()
	}()

	go func() {
		handler.ServeHTTP(rw2, r2)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if rw.Code == http.StatusOK && rw2.Code == http.StatusOK {
		t.Fatalf("One request should have been busy, request 1: %v, request 2: %v", rw.Code, rw2.Code)
	}
}

func TestReleasesConnectionLockWhenFinished(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	handler := NewLimitHandler(1, newTestHandler(ctx))
	rw, r := setup(ctx)
	rw2, r2 := setup(ctx2)

	cancel()
	cancel2()

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		handler.ServeHTTP(rw, r)
		waitGroup.Done()
		handler.ServeHTTP(rw2, r2)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	if rw.Code != http.StatusOK || rw2.Code != http.StatusOK {
		t.Fatalf("One request should have been busy, request 1: %v, request 2: %v", rw.Code, rw2.Code)
	}
}
