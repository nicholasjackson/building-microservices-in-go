package jwt

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

var rsaPrivate *rsa.PrivateKey
var rsaPublic *rsa.PublicKey

func init() {
	bytes, err := ioutil.ReadFile("./sample_key.priv")
	if err != nil {
		log.Fatal("Unable to read private key", err)
	}

	rsaPrivate, err = crypto.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		log.Fatal("Unable to parse private key", err)
	}

	bytes, err = ioutil.ReadFile("./sample_key.pub")
	if err != nil {
		log.Fatal("Unable to read public key", err)
	}

	rsaPublic, err = crypto.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		log.Fatal("Unable to parse public key", err)
	}
}

// GenerateJWT creates a new JWT and signs it with the private key
func GenerateJWT() []byte {
	claims := jws.Claims{}
	claims.SetExpiration(time.Now().Add(time.Duration(1440*3650) * time.Minute))
	claims.Set("userID", "abcsd232jfjf")
	claims.Set("accessLevel", "user")

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	b, _ := jwt.Serialize(rsaPrivate)

	return b
}

// ValidateJWT validates that the given slice is a valid JWT and the signature matches
// the public key
func ValidateJWT(token []byte) error {
	jwt, err := jws.ParseJWT(token)
	if err != nil {
		return fmt.Errorf("Unable to parse token: %v", err)
	}

	if err = jwt.Validate(rsaPublic, crypto.SigningMethodRS256); err != nil {
		return fmt.Errorf("Unable to validate token: %v", err)
	}

	return nil
}
