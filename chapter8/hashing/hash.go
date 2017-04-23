package hashing

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"math/big"

	"github.com/nicholasjackson/building-microservices-in-go/chapter8/utils"
)

// Hash is a structure which is capable of generating and comparing sha512
// hashes derived from a string
type Hash struct {
	peppers []string
}

// New creates a new hash and seeds with a list of peppers
func New(peppers []string) *Hash {
	return &Hash{peppers: peppers}
}

// GenerateHash hashes the input string using sha512 with a salt and pepper.
// returns the hash and the salt
func (h *Hash) GenerateHash(input string, withSalt bool, withPepper bool) (hash string, salt string) {
	pepper := ""

	if withPepper {
		pepper = h.getRandomPepper()
	}

	if withSalt {
		salt = GenerateRandomSalt()
	}

	hash = h.createHash(input, salt, pepper)

	return
}

// Compare checks the string against the salted hash, will loop through
// all peppers until finding success.
func (h *Hash) Compare(input string, salt string, withPepper bool, hash string) bool {

	if withPepper {
		for _, pepper := range h.peppers {
			created := h.createHash(input, salt, pepper)
			if created == hash {
				return true
			}
		}
	} else {
		return h.createHash(input, salt, "") == hash
	}

	return false
}

// GenerateRandomSalt generates a random string 32 bytes in length
func GenerateRandomSalt() string {
	return utils.GenerateRandomString(32)
}

func (h *Hash) createHash(input string, salt string, pepper string) string {
	stringToHash := salt + pepper + input

	sha := sha512.New()
	sha.Write([]byte(stringToHash))

	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

func (h *Hash) getRandomPepper() string {
	max := big.NewInt(int64(len(h.peppers)))
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	return h.peppers[r.Int64()]
}
