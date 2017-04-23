package crypto

// SignMessage signs a message with a private key and returns
// the signature as a string.
func SignMessage(message string, privateKey []byte) (string, error) {

	return "", nil
}

// ValidateSignature check that the given message and signature
// correspond with the given public key.
func ValidateSignature(message string, signature string, pubicKey []byte) (bool, error) {

	return false, nil
}
