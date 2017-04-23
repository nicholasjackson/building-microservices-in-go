package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	// BlockPublicKey contains the value which defines a Public Key in a PEM block
	BlockPublicKey = "PUBLIC KEY"
	// BlockPrivateKey contains the vlaue which defines a PrivateKey in a PEM block
	BlockPrivateKey = "PRIVATE KEY"
)

// UnmarshalRSAPublicKeyFromFile reads the given file and returns a PublicKey
func UnmarshalRSAPublicKeyFromFile(file string) (*rsa.PublicKey, error) {

	block, err := readFileToBlock(file, BlockPublicKey)
	if err != nil {
		return nil, err
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse public key: %s", err)
	}

	return pub.(*rsa.PublicKey), nil
}

// UnmarshalRSAPrivateKeyFromFile reads the given file and returns a PrivateKey
func UnmarshalRSAPrivateKeyFromFile(file string) (*rsa.PrivateKey, error) {

	block, err := readFileToBlock(file, BlockPrivateKey)
	if err != nil {
		return nil, err
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse private key: %s", err)
	}

	return priv, nil
}

func readFileToBlock(file string, blockType string) (*pem.Block, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to open keyfile: %s", err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("Unable to read data from file: %s", err)
	}

	block, _ := pem.Decode(bytes)
	if block == nil && block.Type != blockType {
		return nil, fmt.Errorf("Unable to decode public key")
	}

	return block, nil
}

// GenerateRandomString with crypto/rand
func GenerateRandomString(length int) string {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return ""
	}

	return fmt.Sprintf("%x", k)
}
