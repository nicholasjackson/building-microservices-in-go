package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/nicholasjackson/building-microservices-in-go/chapter8/symmetric"
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

// EncryptDataWithPublicKey encrypts the given data with the public key
func EncryptDataWithPublicKey(data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublic, data, nil)
}

// DecryptDataWithPrivateKey decrypts the given data with the private key
func DecryptDataWithPrivateKey(data []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPrivate, data, nil)
}

// EncryptMessageWithPublicKey encrypts the given string and retuns the encrypted
// result base64 encoded
func EncryptMessageWithPublicKey(message string) (string, error) {

	modulus := rsaPublic.N.BitLen() / 8
	hashLength := 256 / 4
	maxLength := modulus - (hashLength * 2) - 2

	if len(message) > maxLength {
		return "", fmt.Errorf("The maximum message size must not exceed: %d", maxLength)
	}

	data, err := EncryptDataWithPublicKey([]byte(message))
	return base64.StdEncoding.EncodeToString(data), err
}

// DecryptMessageWithPrivateKey decrypts the given base64 encoded ciphertext with
// the private key and returns plain text
func DecryptMessageWithPrivateKey(message string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}

	data, err = DecryptDataWithPrivateKey(data)
	return string(data), err
}

// EncryptLargeMessageWithPublicKey encrypts the given message by randomly generating
// a cipher.
// Returns the ciphertext for the given message base64 encoded and the key
// used to encypt the message which is encrypted with the public key
func EncryptLargeMessageWithPublicKey(message string) (ciphertext string, cipherkey string, err error) {
	key := utils.GenerateRandomString(16) // 16, 24, 32 keysize, random string is 2 bytes per char so 16 chars returns 32 bytes
	cipherData, err := symmetric.EncryptData([]byte(message), []byte(key))
	if err != nil {
		return "", "", err
	}

	cipherkey, err = EncryptMessageWithPublicKey(key)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(cipherData), cipherkey, nil
}

// DecryptLargeMessageWithPrivateKey decrypts the given base64 encoded message by
// decrypting the base64 encoded key with the rsa private key and then using
// the result to decrupt the ciphertext
func DecryptLargeMessageWithPrivateKey(message, key string) (string, error) {
	keystring, err := DecryptMessageWithPrivateKey(key)
	if err != nil {
		return "", fmt.Errorf("Unable to decrypt key with private key: %s", err)
	}

	messageData, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}

	data, err := symmetric.DecryptData(messageData, []byte(keystring))

	return string(data), err
}
