package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// EncryptData encrypts the given data using AES with the given key
func EncryptData(data []byte, key []byte) ([]byte, error) {
	if err := validateKey(key); err != nil {
		return make([]byte, 0), err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return make([]byte, 0), err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return make([]byte, 0), err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return make([]byte, 0), err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// DecryptData decrypts the given data with the given key
func DecryptData(data []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return make([]byte, 0), err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return make([]byte, 0), err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return make([]byte, 0), fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func validateKey(key []byte) error {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return fmt.Errorf("Invalid key length, keys should be 16, 24, or 32 bytes in length")
	}

	return nil
}
