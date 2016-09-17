package bench

import (
	"fmt"
	"time"
)

// Bench is the main application object it allows the profiling of calls to a remote endpoint defined by a Request
type Bench struct {
	threads  int
	testTime time.Duration
	timeout  time.Duration
	request  Request
}

// Results represents a set of benchmark results
type Results struct {
	TotalRequests  int
	AvgRequestTime time.Duration
	TotalFailures  int
}

// String implements String from the Stringer interface and
// allows results to be serialized to a sting
func (r Results) String() string {

	rStr := fmt.Sprintf("Total Requests: %v\n", r.TotalRequests)
	rStr = fmt.Sprintf("%vAvg Request Time: %v\n", rStr, r.AvgRequestTime)
	return fmt.Sprintf("%vTotal Failures: %v\n", rStr, r.TotalFailures)
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
func (b *Bench) RunBenchmarks() Results {
	r := Results{
		TotalRequests:  0,
		TotalFailures:  0,
		AvgRequestTime: 0,
	}

	startTime := time.Now()
	endTime := startTime.Add(b.testTime)
	totalRequestTime := 0 * time.Second

	for run := true; run; run = (time.Now().Before(endTime)) {
		r.TotalRequests += 1
		requestStart := time.Now()
		err := b.request.Do()
		duration := time.Now().Sub(requestStart)

		if err != nil {
			r.TotalFailures += 1
		} else {
			totalRequestTime += duration
		}
	}

	avgTime := int64(totalRequestTime) / int64(r.TotalRequests-r.TotalFailures)
	r.AvgRequestTime = time.Duration(avgTime)

	return r
}
