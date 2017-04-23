package asymetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

const testMessage = "Lorem ipsum dolor sit amet"

const ciphertext = "uAIwFleoavyAurRCOA+icpYYv709zNHfu9cXh1JMzhQ0gneWw+ncWYTiH2Z6FlhTqcOiMq+A1LtlVyP0bQo3PoMegCqi0gFdE4+oB6KLyCFvofMnzJ5vsQedx3ImipQqVmdr5h1MeqiS5EK0vfvdn3e1KszCktrK/aYWwBKql3yY22wPxKn3DvwzgxjUE/VuOYwO0o8sGzTbAxSrXTp5BBaCRvhyrSdQBg6/s6ozrsOfQMhjdlq68Gwdw1DBIxVopmeJYmrd8Aj4yrKTD1Eqt3fBJbHW3Qp6lc+iLRT2wmkP+GeQ7Y0sZjOQ8hHdGz1Hqour/CQ1cPMGWchHzg/k8w=="

func TestEncryptWithPublicKey(t *testing.T) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublic, []byte(testMessage), nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
}

func TestDecryptWithPrivateKey(t *testing.T) {
	data, _ := base64.StdEncoding.DecodeString(ciphertext)
	plainData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPrivate, data, nil)
	if err != nil {
		t.Fatal(err)
	}

	if string(plainData) != testMessage {
		t.Fatalf("Expecting %v, got %v", testMessage, string(plainData))
	}
}
