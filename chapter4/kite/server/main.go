package main

import (
	"fmt"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("math", "1.0.0")

	// Add our handler method with the name "square"
	k.HandleFunc("Hello", func(r *kite.Request) (interface{}, error) {
		name, _ := r.Args.One().String()

		return fmt.Sprintf("Hello %v", name), nil
	}).DisableAuthentication()

	// Attach to a server with port 3636 and run it
	k.Config.Port = 8091
	k.Run()

}
