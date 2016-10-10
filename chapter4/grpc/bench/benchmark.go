package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
	proto "github.com/nicholasjackson/building-microservices-in-go/chapter4/grpc/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Benchmarking application")

	b := bench.New(400, 300*time.Second, 90*time.Second, 5*time.Second)
	b.AddOutput(301*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(GrpcRequest)
}

// GrpcRequest is executed by benchmarks
func GrpcRequest() error {

	conn, err := grpc.Dial("consul.acet.io:9000", grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := proto.NewKittensClient(conn)
	_, err = client.Hello(context.Background(), &proto.Request{Name: "Nic"})

	if err != nil {
		return err
	}

	return nil
}
