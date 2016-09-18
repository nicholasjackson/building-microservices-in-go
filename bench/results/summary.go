package results

import (
	"fmt"
	"time"

	"github.com/nicholasjackson/building-microservices-in-go/bench/errors"
)

// Results represents a set of benchmark results
type table struct {
	TotalRequests  int
	AvgRequestTime time.Duration
	TotalSuccess   int
	TotalFailures  int
	TotalTimeouts  int
}

// String implements String from the Stringer interface and
// allows results to be serialized to a sting
func (r table) String() string {

	rStr := fmt.Sprintf("Total Requests: %v\n", r.TotalRequests)
	rStr = fmt.Sprintf("%vAvg Request Time: %v\n", rStr, r.AvgRequestTime)
	rStr = fmt.Sprintf("%vTotal Success: %v\n", rStr, r.TotalSuccess)
	rStr = fmt.Sprintf("%vTotal Timeouts: %v\n", rStr, r.TotalTimeouts)
	return fmt.Sprintf("%vTotal Failures: %v\n", rStr, r.TotalFailures)
}

// TabularResults processes a set of results and writes a tabular
// summary to the standard output
type TabularResults struct{}

// Process implements the Result interface Process
func (t *TabularResults) Process(results []Result) {

	r := table{
		TotalRequests:  0,
		TotalFailures:  0,
		TotalSuccess:   0,
		TotalTimeouts:  0,
		AvgRequestTime: 0,
	}

	totalRequestTime := 0 * time.Second

	for _, result := range results {
		r.TotalRequests++

		if result.Error != nil {
			if _, ok := result.Error.(errors.Timeout); ok {
				r.TotalTimeouts++
			}

			r.TotalFailures++
		} else {
			r.TotalSuccess++
			totalRequestTime += result.RequestTime
		}
	}

	avgTime := int64(totalRequestTime) / int64(r.TotalRequests-r.TotalFailures)
	r.AvgRequestTime = time.Duration(avgTime)

	fmt.Println(r)
}
