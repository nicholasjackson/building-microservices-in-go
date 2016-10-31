package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eapache/go-resiliency/deadline"
)

func main() {
	switch os.Args[1] {
	case "slow":
		makeNormalRequest()
	case "timeout":
		makeTimeoutRequest()
	}
}

func makeNormalRequest() {
	slowFunction()
}

func makeTimeoutRequest() {
	dl := deadline.New(1 * time.Second)
	err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction()
		return nil
	})

	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("Timeout")
	default:
		fmt.Println(err)
	}
}

func slowFunction() {
	for i := 0; i < 100; i++ {
		fmt.Println("Loop: ", i)
		time.Sleep(1 * time.Second)
	}
}
