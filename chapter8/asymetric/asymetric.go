package asymetric

import (
	"crypto/rsa"
	"log"

	"github.com/nicholasjackson/building-microservices-in-go/chapter8/utils"
)

var rsaPublic *rsa.PublicKey
var rsaPrivate *rsa.PrivateKey

func init() {
	var err error
	rsaPrivate, err = utils.UnmarshalRSAPrivateKeyFromFile("../keys/sample_key.priv")
	if err != nil {
		log.Fatal("Unable to read private key", err)
	}

	rsaPublic, err = utils.UnmarshalRSAPublicKeyFromFile("../keys/sample_key.pub")
	if err != nil {
		log.Fatal("Unable to read public key", err)
	}
}
