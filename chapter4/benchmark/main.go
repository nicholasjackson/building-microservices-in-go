package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/nicholasjackson/building-microservices-in-go/chapter4/benchmark/data"
	"github.com/nicholasjackson/building-microservices-in-go/chapter4/benchmark/handlers"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to file")
var store *data.MongoStore

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		fmt.Println("Running with CPU profile")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Finished")
		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal(err)
			}
			runtime.GC()
			pprof.WriteHeapProfile(f)
			defer f.Close()
		}
		if *cpuprofile != "" {
			pprof.StopCPUProfile()
		}

		os.Exit(0)
	}()

	store = waitForDB()
	clearDB()
	setupData()

	handler := handlers.Search{DataStore: store}
	err := http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Exit")
}

func waitForDB() *data.MongoStore {
	for i := 0; i < 10; i++ {
		store, err := data.NewMongoStore("mongodb")
		if err == nil {
			return store
		}

		fmt.Println("Waiting for DB Connection")
		time.Sleep(1 * time.Second)
	}

	return nil
}

func clearDB() {
	store.DeleteAllKittens()
}

func setupData() {
	store.InsertKittens(
		[]data.Kitten{
			data.Kitten{
				Id:     "1",
				Name:   "Felix",
				Weight: 12.3,
			},
			data.Kitten{
				Id:     "2",
				Name:   "Fat Freddy's Cat",
				Weight: 20.0,
			},
			data.Kitten{
				Id:     "3",
				Name:   "Garfield",
				Weight: 35.0,
			},
		})
}
