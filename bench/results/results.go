package results

import "time"

// Result is a structure that encapsulates processable result data
type Result struct {
	Timestamp   time.Time
	RequestTime time.Duration
	Error       error
}

// Results is an interface which result processors must implement
type Results interface {
	// Process the result set
	Process(results []Result)
}
