package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
	proto "github.com/nicholasjackson/building-microservices-in-go/chapter6/grpc/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func main() {
	conn, _ = grpc.Dial("consul.acet.io:9000", grpc.WithInsecure())
	defer conn.Close()

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
	client := proto.NewKittensClient(conn)
	_, err := client.Hello(context.TODO(), &proto.Request{Name: "Nic"})

	if err != nil {
		return err
	}

	return nil
}
