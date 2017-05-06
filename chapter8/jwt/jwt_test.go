package jwt

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

const validJWT = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NMZXZlbCI6InVzZXIiLCJleHAiOjE4MDc3MDEyNDYsInVzZXJJRCI6ImFiY3NkMjMyamZqZiJ9.iQxUbQuEy4Jh4oTkkz0OPGvS86xOWJjdzxHHDBeAolv0982pXKPBMWskSJDF3F8zd8a8nMIlQ5m9tzePoJWe_E5B9PRJEvYAUuSp6bGm7-IQEum8EzHq2tMvYtPl9uzXgOU4C_pIjZh5CjFUeZLk5tWKwOOo8pW4NUSxsV2ZRQ_CGfIrBqEQgKRodeLTcQ4wJkLBILBzmAqTVl-5sLgBEoZ76C_gcvS6l5HAwEAhmiCqtDMX46o8pA72Oa6NiVRsgxrhrKX9rDUBdJAxNwFAwCjTv6su0jTZvkYD80Li9aXiMuM9NX7q5gncbEhfko_byTYryLsmmaUSXNBlnvC_nQ"

func TestGenerateJWT(t *testing.T) {
	b := GenerateJWT()

	fmt.Println(string(b))
}

func TestValidateJWT(t *testing.T) {
	err := ValidateJWT([]byte(validJWT))

	if err != nil {
		t.Fatal(err)
	}
}

func TestSplitIntoDataAndSignature(t *testing.T) {
	data := strings.Split(validJWT, ".")

	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(strings.Join(data[:2], "."))

	file, err = os.OpenFile("signature.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(data[2])

	signature, err := base64.RawURLEncoding.DecodeString(data[2])
	if err != nil {
		log.Fatal(err)
	}

	file2, err := os.OpenFile("signature.sha256", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	file2.Write(signature)
}
