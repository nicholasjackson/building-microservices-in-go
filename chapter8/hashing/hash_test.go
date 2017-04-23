package hashing

import (
	"fmt"
	"testing"
)

const (
	original              = "HelloWorld1"
	saltedAndPepperedSalt = "56c01edd9608037d4d262ebb6186b5e092c780a7edca161b43818f346d45c7ad"
	saltedAndPepperedHash = "fe2bb81c63ef3187dd5ef0d92f474d4483f2c11ca8c4c4c5fc3a9e1c56b32b5e946ad6bb95b50381a87b7ea295b48ff9999ff71023399d90463d5d011ea6fd57"
	saltedHash            = "6e580ce55ddbb1e7d6d616b6c8480a4b90875c459ce7a0b59f30ace96995e3b11173c0456754b73827fbd27e9d15c52772e2f409b87003321bce6be97852c2ec"
	saltedSalt            = "2cc9ae49bd3fd3d223b5de3b40efb0c0ebcd3387117f0d8a3b680f286ece94a9"
	plainHash             = "a4db351d57adf4b71105ef2b13138ea50d539c93a83b471974a7a7c1f8b132cd267e11266529eaaf08f05e516dbebf03133688826cb538eebc626bb06ad1ebd8"
)

var peppers = []string{
	"47278c6cd6353a278a2a5929f77752ac429acd59cbded92cdf88a68fdfb9ac2f",
	"b0aa0db641509c907459bf95445a78413dc310b2fdd2d8962562d8e3327a04e0",
	"90ffbdb56950f4be181f752d8fe2f9dad682ef22fc163429ebe288bbc0a91804",
	"189f58d27aa9e979e685acdb10e8587a09e1bcae4ec93bea05522f4e5b32f3b5",
	"e1e9c4e5bcafc552ded0c849fbc896bd6fa7c03c0410b60d1e7541832be66fa6",
}

func BenchmarkGenerateHashWithSaltAndPepper(b *testing.B) {
	h := New(peppers)

	for i := 0; i < b.N; i++ {
		_, _ = h.GenerateHash(original, true, true)
	}
}

func BenchmarkCompareHash5Peppers(b *testing.B) { benchmarkCompare(5, b) }

func BenchmarkCompareHash10Peppers(b *testing.B) { benchmarkCompare(10, b) }

func BenchmarkCompareHash100Peppers(b *testing.B) { benchmarkCompare(100, b) }

func BenchmarkCompareHash1000Peppers(b *testing.B) { benchmarkCompare(1000, b) }

func benchmarkCompare(pepperSize int, b *testing.B) {
	p := generatePeppers(pepperSize)
	h := New(p)
	hash, salt := h.GenerateHash(original, true, true)

	for i := 0; i < b.N; i++ {
		h.Compare(original, salt, true, hash)
	}
}

func BenchmarkGeneratePlainHash(b *testing.B) {
	h := New(peppers)

	for i := 0; i < b.N; i++ {
		_, _ = h.GenerateHash(original, false, false)
	}
}

func BenchmarkCompareSaltedHash(b *testing.B) {
	h := New(peppers)
	hash, salt := h.GenerateHash(original, true, false)

	for i := 0; i < b.N; i++ {
		h.Compare(original, salt, false, hash)
	}
}

func BenchmarkComparePlainHash(b *testing.B) {
	h := New(peppers)
	hash, _ := h.GenerateHash(original, false, false)

	for i := 0; i < b.N; i++ {
		h.Compare(original, "", false, hash)
	}
}

func TestGenerateHash(t *testing.T) {
	h := New(peppers)
	hash, salt := h.GenerateHash(original, true, true)

	fmt.Println("Salted and Peppered Hash")
	fmt.Println("Salt: ", salt)
	fmt.Println("Hash: ", hash)
	fmt.Println("-----")
}

func TestGenerateSaltedHash(t *testing.T) {
	h := New(peppers)
	hash, salt := h.GenerateHash(original, true, false)

	fmt.Println("Salted Hash")
	fmt.Println("Hash: ", hash)
	fmt.Println("Salt: ", salt)
	fmt.Println("-----")
}

func TestGeneratePlainHash(t *testing.T) {
	h := New(peppers)
	hash, _ := h.GenerateHash(original, false, false)

	fmt.Println("Plain Hash")
	fmt.Println("Hash: ", hash)
	fmt.Println("-----")
}

func TestCompareSaltedAndPeppered(t *testing.T) {
	h := New(peppers)
	success := h.Compare(original, saltedAndPepperedSalt, true, saltedAndPepperedHash)

	if !success {
		t.Fatal("Should have successfully compared to original")
	}
}

func TestComparesSaltedCorrectly(t *testing.T) {
	h := New(peppers)
	success := h.Compare(original, saltedSalt, false, saltedHash)

	if !success {
		t.Fatal("Should have successfully compared to original")
	}
}

func TestComparesPlainCorrectly(t *testing.T) {
	h := New(peppers)
	success := h.Compare(original, "", false, plainHash)

	if !success {
		t.Fatal("Should have successfully compared to original")
	}
}

func generatePeppers(n int) []string {
	p := make([]string, 0)

	for i := 0; i < n; i++ {
		h := GenerateRandomSalt()
		p = append(p, h)
	}

	return p
}
