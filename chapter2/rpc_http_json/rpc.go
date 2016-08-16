package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

//curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:1234

const port = 1234

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

func main() {
	if os.Args[1] == "server" {
		server()
	}
}

func server() {
	helloWorld := new(HelloWorldHandler)
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
		err := rpc.ServeRequest(serverCodec)
		if err != nil {
			log.Printf("Error while serving JSON request: %v", err)
			http.Error(w, "Error while serving JSON request, details have been logged.", 500)
			return
		}
	}))
}

type HelloWorldRequest struct {
	Name string
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *HelloWorldRequest, reply *HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
