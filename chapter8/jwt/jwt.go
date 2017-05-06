package jwt

import (
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/nicholasjackson/building-microservices-in-go/chapter8/utils"
)

var rsaPrivate *rsa.PrivateKey
var rsaPublic *rsa.PublicKey

func init() {
	var err error
	rsaPrivate, err = utils.UnmarshalRSAPrivateKeyFromFile("../keys/sample_key.priv")
	if err != nil {
		log.Fatal("Unable to parse private key", err)
	}

	rsaPublic, err = utils.UnmarshalRSAPublicKeyFromFile("../keys/sample_key.pub")
	if err != nil {
		log.Fatal("Unable to parse public key", err)
	}
}

// GenerateJWT creates a new JWT and signs it with the private key
func GenerateJWT() []byte {
	claims := jws.Claims{}
	claims.SetExpiration(time.Now().Add(2880 * time.Minute))
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
