package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	rootKey, rootCert := generateRootCA()
	_ = generateIntermediateCert(rootCert, rootKey)
}

func generateRootCA() (*ecdsa.PrivateKey, *x509.Certificate) {
	priv, err := generatePrivateKey()
	if err != nil {
		log.Fatal("Unable to generate private key: ", err)
	}

	err = savePrivateKey(priv, "root_key.pem")
	if err != nil {
		log.Fatal("Unable to save private key: ", err)
	}

	data, err := generateRootX509Certificate(priv)
	if err != nil {
		log.Fatal("Unable to generate certificate: ", err)
	}

	err = saveX509Certificate(data, "root_cert.pem")
	if err != nil {
		log.Fatal("Unable to save cert: ", err)
	}

	cert, err := x509.ParseCertificate(data)
	if err != nil {
		log.Fatal("Unable to parse certificiate: ", err)
	}

	return priv, cert
}

func generateIntermediateCert(parentCert *x509.Certificate, parentKey *ecdsa.PrivateKey) *ecdsa.PrivateKey {
	priv, err := generatePrivateKey()
	if err != nil {
		log.Fatal("Unable to generate private key: ", err)
	}

	err = savePrivateKey(priv, "intermediate_key.pem")
	if err != nil {
		log.Fatal("Unable to save private key: ", err)
	}

	data, err := generateIntermediateX509Certificate(priv, parentCert, parentKey)
	if err != nil {
		log.Fatal("Unable to generate certificate: ", err)
	}

	err = saveX509Certificate(data, "intermediate_cert.pem")
	if err != nil {
		log.Fatal("Unable to save cert: ", err)
	}
	return priv
}

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func savePrivateKey(key *ecdsa.PrivateKey, path string) error {
	b, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return fmt.Errorf("Unable to marshal private key: %v", err)
	}

	block := &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}

	keyOut, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open key.pem for writing: %v", err)
	}

	pem.Encode(keyOut, block)
	keyOut.Close()

	return nil
}

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

func generateRootX509Certificate(key *ecdsa.PrivateKey) ([]byte, error) {

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %s", err)
	}

	subjectKey, err := getSubjectKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to get subject key: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Country:            []string{"UK"},
			Organization:       []string{"Acme Co"},
			OrganizationalUnit: []string{"Tech"},
			CommonName:         "RootCert",
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		IsCA:         true,
		SubjectKeyId: subjectKey,

		KeyUsage: x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	return x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
}

func generateIntermediateX509Certificate(
	key *ecdsa.PrivateKey,
	parentCert *x509.Certificate,
	parentKey *ecdsa.PrivateKey) ([]byte, error) {

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %s", err)
	}

	subjectKey, err := getSubjectKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to get public subject key")
	}

	authKey, err := getSubjectKey(parentKey)
	if err != nil {
		return nil, fmt.Errorf("unable to get public subject key")
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Country:            []string{"UK"},
			Organization:       []string{"Acme Co"},
			OrganizationalUnit: []string{"Tech"},
			CommonName:         "Intermediate",
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		SubjectKeyId:   subjectKey,
		AuthorityKeyId: authKey,

		KeyUsage: x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature |
			x509.KeyUsageCertSign |
			x509.KeyUsageCRLSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	return x509.CreateCertificate(rand.Reader, &template, parentCert, &key.PublicKey, parentKey)
}

func getSubjectKey(key *ecdsa.PrivateKey) ([]byte, error) {
	publicKey, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %s", err)
	}

	var subPKI subjectPublicKeyInfo
	_, err = asn1.Unmarshal(publicKey, &subPKI)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %s", err)
	}

	return bigIntHash(subPKI.SubjectPublicKey.Bytes), nil
}

func bigIntHash(n []byte) []byte {
	h := sha1.New()
	h.Write(n)
	return h.Sum(nil)
}

func saveX509Certificate(data []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to open cert.pem for writing: %s", err)
	}

	block := &pem.Block{Type: "CERTIFICATE", Bytes: data}
	pem.Encode(file, block)
	file.Close()

	return nil
}
