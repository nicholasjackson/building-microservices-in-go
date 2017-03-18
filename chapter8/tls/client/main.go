package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	rootCert, err := ioutil.ReadFile("../generate_keys/root_cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	applicationCert, err := ioutil.ReadFile("../generate_keys/application_cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootCert)
	if !ok {
		panic("failed to parse root certificate")
	}
	ok = roots.AppendCertsFromPEM(applicationCert)
	if !ok {
		panic("failed to parse root certificate")
	}

	tlsConf := &tls.Config{RootCAs: roots}

	tr := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: tr}
	//	client := http.DefaultClient

	resp, err := client.Get("https://localhost:8433")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
