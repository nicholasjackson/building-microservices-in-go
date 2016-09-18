package main

import (
	"fmt"
	"time"

	"github.com/nicholasjackson/building-microservices-in-go/bench"
	"github.com/nicholasjackson/building-microservices-in-go/bench/results"
)

func main() {
	fmt.Println("Benchmarking application")

	b := bench.New(2, 10*time.Second, 2*time.Second, &GoMicroRequest{})
	r := b.RunBenchmarks()

	fmt.Println("Benchmarking completed")

	fmt.Print("\nResults:\n")

	table := results.TabularResults{}
	table.Process(r)

	plot := results.PlotResults{}
	plot.Process(r)
}

// GoMicroRequest is an example request object
type GoMicroRequest struct {
}

// Do is an example implementation of a request
func (r *GoMicroRequest) Do() error {

	if time.Now().UnixNano()%10 == 0 {
		time.Sleep(3 * time.Second)
		return fmt.Errorf("Timeout")
	}

	if time.Now().UnixNano()%2 == 0 {
		return fmt.Errorf("Just an error")
	}

	return nil
}
