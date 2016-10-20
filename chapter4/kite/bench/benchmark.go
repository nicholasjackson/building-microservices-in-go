package main

import (
	"fmt"
	"time"
	"os"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
	"github.com/koding/kite"
	"github.com/nicholasjackson/bench"
)

var kittenServer *kite.Client

func main() {
	fmt.Println("Benchmarking application")

	k := kite.New("exp2", "1.0.0")

	kittenServer = k.NewClient("http://consul.acet.io:8091/kite")
	kittenServer.Dial()

	b := bench.New(400, 300*time.Second, 90*time.Second, 5*time.Second)
	b.AddOutput(0*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(KiteRequest)
}

// GoMicroRequest is executed by benchmarks
func KiteRequest() error {
	_, err := kittenServer.Tell("Hello", "Nic") // call "square" method with argument 4
	if err != nil {
		return err
	}

	return nil
}
