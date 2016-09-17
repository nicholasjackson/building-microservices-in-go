package main

import (
	"fmt"
	"time"

	"github.com/nicholasjackson/building-microservices-in-go/bench"
)

func main() {
	fmt.Println("Benchmarking application")

	b := bench.New(10, 1*time.Second, 2*time.Second, &GoMicroRequest{})
	r := b.RunBenchmarks()

	fmt.Println("Benchmarking completed")

	fmt.Print("\nResults:\n")
	fmt.Println(r)
}

type GoMicroRequest struct {
	err error
}

func (r *GoMicroRequest) Do() error {
	if r.err != nil {
		r.err = nil
	} else {
		r.err = fmt.Errorf("sd")
	}
	return r.err
}
