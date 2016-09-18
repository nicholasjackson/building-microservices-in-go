package bench

import (
	"fmt"
	"time"

	"github.com/nicholasjackson/building-microservices-in-go/bench/errors"
	"github.com/nicholasjackson/building-microservices-in-go/bench/results"
)

// Bench is the main application object it allows the profiling of calls to a remote endpoint defined by a Request
type Bench struct {
	threads  int
	testTime time.Duration
	timeout  time.Duration
	request  Request
}

// New creates a new bench and intializes the intial values of...
// threads: the number of concurrent threads to execute
// testDuration: the duration to run the benchmarks for
// timeout: the timeout value
// request: the Request to execute
//
// returns a new Bench instance
func New(
	threads int,
	testDuration time.Duration,
	timeout time.Duration,
	request Request) *Bench {

	return &Bench{
		threads:  threads,
		testTime: testDuration,
		timeout:  timeout,
		request:  request,
	}
}

// RunBenchmarks executes the benchmarks based upon the given criteria
//
// Returns a resultset
func (b *Bench) RunBenchmarks() []results.Result {

	startTime := time.Now()
	endTime := startTime.Add(b.testTime)

	semaphore := make(chan struct{}, b.threads)
	abandon := make(chan bool, b.threads) //  use a buffered channel as if we don't when nothing reads we block
	out := make(chan results.Result)
	var results []results.Result

	// handle the returns from the threads
	go handleResult(out, &results)

	for run := true; run; run = (time.Now().Before(endTime)) {

		semaphore <- struct{}{} // blocks when channel is full

		// execute a request
		go doRequest(b.request, b.timeout, semaphore, out, abandon)
	}

	// Instruct all threads to abandon
	for i := 0; i < b.threads; i++ {
		abandon <- true
	}

	// wait for threads to return
	// once we fill the buffer to capacity we know that
	// we can safely quit
	for i := 0; i < cap(semaphore); i++ {
		semaphore <- struct{}{}
	}

	// close the output channel so the handle result method returns
	close(out)

	return results
}

func handleResult(out chan results.Result, results *[]results.Result) {

	for result := range out {
		*results = append(*results, result)
	}
}

func doRequest(request Request, timeout time.Duration, semaphore chan struct{}, out chan results.Result, abandon chan bool) {

	defer func() {
		<-semaphore // notify we are done at the end of the routine
	}()

	requestStart := time.Now()

	timeoutC := time.After(timeout)
	complete := make(chan results.Result)

	go func() {
		err := request.Do()
		complete <- results.Result{
			Error:       err,
			RequestTime: time.Now().Sub(requestStart),
			Timestamp:   time.Now(),
		}
	}()

	var ret results.Result

	select {
	case ret = <-complete:
	case <-timeoutC:
		ret.Error = errors.Timeout{Message: "Timeout error"}
		ret.RequestTime = timeout
		ret.Timestamp = time.Now()
	case <-abandon:
		fmt.Println("Abandon")
	}

	out <- ret
}
