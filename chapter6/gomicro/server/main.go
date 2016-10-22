package main

import (
	"context"
	"flag"
	"os"
	"runtime/pprof"

	log "github.com/golang/glog"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	kittens "github.com/nicholasjackson/building-microservices-in-go/chapter6/gomicro/proto"
)

type Kittens struct{}

func (s *Kittens) Hello(ctx context.Context, req *kittens.Request, rsp *kittens.Response) error {
	rsp.Msg = server.DefaultId + ": Hello " + req.Name

	return nil
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	cmd.Init()

	server.Init(
		server.Name("bmigo.micro.Kittens"),
		server.Version("1.0.0"),
		server.Address(":8091"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(Kittens),
		),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
