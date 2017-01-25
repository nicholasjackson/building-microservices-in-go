package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var method = flag.String("method", "hash", "Specify hash to hash an input string, check to check a hash against a string")
var input = flag.String("input", "", "The input string to hash")
var salt = flag.String("salt", "", "The salt to use with the string")
var hash = flag.String("hash", "", "The hash to compare to the input")

var peppers = []string{
	"anvdfkljslkfjasdfklj4123413123",
	"andrfjkasfkljasdfkl341231238979oiuklja.,a",
	"lkjklujiuwekj,.1m318979812l3kj,.mc,.m",
	"cvvcv,mkjkllweioiwoeioooooqwepoi,ljma,l;sfjakl;fj1123123lkj",
	"sdfjkwuoiiweoiqoeikasd,ljms,.vmxc.,vmsdklfjsdklruweioru",
}

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	flag.Parse()

	switch *method {
	case "hash":
		h := createHash(*input, *salt, getRandomPepper())
		fmt.Println(h)
	case "compare":
		if compare(*input, *salt, *hash) {
			fmt.Println("Input equal to hash")
		} else {
			fmt.Println("Input not equal to hash")
		}
	}

}

func createHash(input string, salt string, pepper string) string {
	stringToHash := salt + pepper + input

	sha := sha512.New()
	sha.Write([]byte(stringToHash))

	h := sha.Sum(nil)

	return fmt.Sprintf("%x", h)
}

func compare(input string, salt string, hash string) bool {
	for _, pepper := range peppers {
		h := createHash(input, salt, pepper)
		if h == hash {
			return true
		}
	}

	return false
}

func getRandomPepper() string {
	r := random.Intn(len(peppers))
	return peppers[r]
}
