package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
)

func main() {

	fmt.Println("Benchmarking application")

	b := bench.New(1, 37*time.Second, 0*time.Second, 2*time.Second)
	b.AddOutput(0*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.AddOutput(0*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.RunBenchmarks(AmazonRequest)
}

// AmazonRequest is an example to benchmark a call to googles homepage
func AmazonRequest() error {

	resp, err := http.Get("http://www.amazon.co.uk/")
	defer func(response *http.Response) {
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}(resp)

	if err != nil || resp.StatusCode != 200 {
		return err
	}

	return nil
}
