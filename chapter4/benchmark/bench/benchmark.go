package main

import (
	"bytes"
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

	b := bench.New(400, 300*time.Second, 90*time.Second, 5*time.Second)
	b.AddOutput(301*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(Request)
}

// Request is executed by benchmarks
func Request() error {

	req, err := http.NewRequest(
		"POST",
		"http://consul.acet.io:8323/search",
		bytes.NewBuffer([]byte(`{"query":"Fat Freddy's Cat"}`)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed with status: %v", resp.Status)
	}

	return nil
}
