package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/micro/go-micro/client/rpc"
	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
	kittens "github.com/nicholasjackson/building-microservices-in-go/chapter4/gomicro/proto"
)

func main() {
	fmt.Println("Benchmarking application")

	b := bench.New(20, 30*time.Second, 15*time.Second, 2*time.Second)
	b.AddOutput(31*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(GoMicroRequest)
}

// GoMicroRequest is executed by benchmarks
func GoMicroRequest() error {

	c := rpc.NewClient()
	request := c.NewRequest("bmigo.micro.Kittens", "Kittens.List", &kittens.Request{Name: "Nic"})
	response := &kittens.Response{}

	err := c.CallRemote(
		context.Background(),
		"consul.acet.io:8091",
		request,
		response)

	if err != nil {
		return err
	}

	return nil
}
