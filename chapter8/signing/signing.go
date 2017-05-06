package signing

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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

func SignMessageWithPrivateKey(message string) (string, error) {
	rnd := rand.Reader
	hashed := sha256.Sum256([]byte(message))

	signature, err := rsa.SignPKCS1v15(rnd, rsaPrivate, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("error from signing: %s", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func ValidateMessageWithPublicKey(message string, signature string) error {
	hashed := sha256.Sum256([]byte(message))

	signatureData, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("unable to decode signature: %s", err)
	}

	err = rsa.VerifyPKCS1v15(rsaPublic, crypto.SHA256, hashed[:], signatureData)
	if err != nil {
		return fmt.Errorf("unable to validate signature: %s", err)
	}

	return nil
}
