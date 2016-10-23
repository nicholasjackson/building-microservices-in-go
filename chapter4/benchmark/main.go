package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"runtime/pprof"

	"github.com/nicholasjackson/building-microservices-in-go/chapter4/benchmark/data"
	"github.com/nicholasjackson/building-microservices-in-go/chapter4/benchmark/handlers"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to file")

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

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		defer f.Close()
	}

	store, err := data.NewMongoStore("mongodb")
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Search{DataStore: store}
	err = http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
