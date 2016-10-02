package main

import (
	"fmt"
	"os"
	"time"

	"github.com/koding/kite"
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
	b.RunBenchmarks(GoMicroRequest)
}

// GoMicroRequest is executed by benchmarks
func GoMicroRequest() error {

	k := kite.New("exp2", "1.0.0")

	// Connect to our math kite
	kittenServer := k.NewClient("http://consul.acet.io:8091/kite")
	kittenServer.Dial()

	_, err := kittenServer.Tell("Hello", "Nic") // call "square" method with argument 4

	if err != nil {
		return err
	}

	return nil
}
