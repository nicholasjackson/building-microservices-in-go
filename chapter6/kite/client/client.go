package main

import (
	"fmt"
	"log"

	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"github.com/koding/kite/protocol"
)

func main() {
	k := kite.New("exp2", "1.0.0")

	c := config.MustGet()
	k.Config = c
	k.SetLogLevel(kite.DEBUG)

	kites, err := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "math",
	})
	if err != nil {
		log.Fatalln(err)
	}

	kittenServer := kites[0]
	connected, err := kittenServer.DialForever()
	if err != nil {
		k.Log.Fatal(err.Error())
	}

	// Wait until connected
	<-connected

	response, _ := kittenServer.Tell("Hello", "Nic") // call "square" method with argument 4
	message, _ := response.String()
	fmt.Println("result:", message)
}
