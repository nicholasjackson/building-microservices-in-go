package main

import (
	"fmt"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("exp2", "1.0.0")

	// Connect to our math kite
	kittenServer := k.NewClient("http://localhost:8091/kite")
	kittenServer.Dial()

	response, _ := kittenServer.Tell("Hello", "Nic") // call "square" method with argument 4
	message, _ := response.String()
	fmt.Println("result:", message)
}
