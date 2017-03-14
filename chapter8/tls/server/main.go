package main

import (
	"fmt"
	"log"
	"net/http"
)

// generate key:
// openssl ecparam -genkey -name secp384r1 -out server.key

// generate cert:
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(rw, "Hello World")
	})

	err := http.ListenAndServeTLS(":8433", "../certs/3_application/certs/www.example.com.cert.pem", "../certs/3_application/private/www.example.com.key.pem", nil)

	log.Fatal(err)
}
