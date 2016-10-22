package main

import (
	"fmt"

	_ "net/http/pprof"
	"net/url"

	"github.com/koding/kite"
	"github.com/koding/kite/config"
)

func main() {

	k := kite.New("math", "1.0.0")
	c := config.MustGet()
	k.Config = c
	k.Config.KontrolURL = "http://kontrol:6000/kite"

	k.RegisterForever(&url.URL{Scheme: "http", Host: "127.0.0.1:8091", Path: "/kite"})

	// Add our handler method with the name "square"
	k.HandleFunc("Hello", func(r *kite.Request) (interface{}, error) {
		name, _ := r.Args.One().String()

		return fmt.Sprintf("Hello %v", name), nil
	}).DisableAuthentication()

	// Attach to a server with port 3636 and run it
	k.Config.Port = 8091
	k.Run()

}
