package main

import (
	"crypto/rand"
	"crypto/rsa"
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

const durationDecade = (time.Hour * 24 * 3650)
const durationYear = (time.Hour * 24 * 365)
const durationMonth = (time.Hour * 24 * 30)

var rootTemplate = x509.Certificate{
	Subject: pkix.Name{
		Country:            []string{"UK"},
		Organization:       []string{"Acme Co"},
		OrganizationalUnit: []string{"Tech"},
		CommonName:         "Root",
	},

	KeyUsage: x509.KeyUsageKeyEncipherment |
		x509.KeyUsageDigitalSignature |
		x509.KeyUsageCertSign |
		x509.KeyUsageCRLSign,
	BasicConstraintsValid: true,
	IsCA: true,
}

var applicationTemplate = x509.Certificate{
	Subject: pkix.Name{
		Country:            []string{"UK"},
		Organization:       []string{"Acme Co"},
		OrganizationalUnit: []string{"Tech"},
		CommonName:         "Application",
	},

	KeyUsage: x509.KeyUsageKeyEncipherment |
		x509.KeyUsageDigitalSignature |
		x509.KeyUsageCertSign |
		x509.KeyUsageCRLSign,
	BasicConstraintsValid: true,
	IsCA: true,
}

var instanceTemplate = x509.Certificate{
	Subject: pkix.Name{
		Country:            []string{"UK"},
		Organization:       []string{"Acme Co"},
		OrganizationalUnit: []string{"Tech"},
		CommonName:         "Instance",
	},

	KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	BasicConstraintsValid: true,
	DNSNames:              []string{"localhost"},
}

func main() {
	rootKey, rootCert := generateKeyAndCertificate(
		&rootTemplate,
		durationDecade,
		"root_key.pem",
		"root_cert.pem",
		"password",
		nil,
		nil)

	applicationKey, applicationCert := generateKeyAndCertificate(
		&applicationTemplate,
		durationYear,
		"application_key.pem",
		"application_cert.pem",
		"password",
		rootKey,
		rootCert)

	_, _ = generateKeyAndCertificate(
		&instanceTemplate,
		durationMonth,
		"instance_key.pem",
		"instance_cert.pem",
		"",
		applicationKey,
		applicationCert)
}

func generateKeyAndCertificate(
	template *x509.Certificate,
	duration time.Duration,
	keyOutputLocation string,
	certficateOutputLocation string,
	password string,
	parentKey *rsa.PrivateKey,
	parentCert *x509.Certificate) (*rsa.PrivateKey, *x509.Certificate) {

	priv := generatePrivateKey()
	savePrivateKey(priv, keyOutputLocation, []byte(password))

	certData := generateX509Certificate(priv, template, duration, parentKey, parentCert)
	saveX509Certificate(certData, certficateOutputLocation)

	cert, err := x509.ParseCertificate(certData)
	if err != nil {
		log.Fatal("Unable to parse certificiate: ", err)
	}

	return priv, cert
}

func generatePrivateKey() *rsa.PrivateKey {
	key, _ := rsa.GenerateKey(rand.Reader, 4096)
	return key
}

func generateX509Certificate(
	key *rsa.PrivateKey,
	template *x509.Certificate,
	duration time.Duration,
	parentKey *rsa.PrivateKey,
	parentCert *x509.Certificate) []byte {

	notBefore := time.Now()
	notAfter := notBefore.Add(duration)

	template.NotBefore = notBefore
	template.NotAfter = notAfter

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		panic(fmt.Errorf("failed to generate serial number: %s", err))
	}

	template.SerialNumber = serialNumber

	subjectKey, err := getSubjectKey(key)
	if err != nil {
		panic(fmt.Errorf("unable to get subject key: %s", err))
	}

	template.SubjectKeyId = subjectKey

	if parentKey == nil {
		parentKey = key
	}

	if parentCert == nil {
		parentCert = template
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, parentCert, &key.PublicKey, parentKey)
	if err != nil {
		panic(err)
	}

	return cert
}

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

func getSubjectKey(key *rsa.PrivateKey) ([]byte, error) {
	publicKey, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %s", err)
	}

	var subPKI subjectPublicKeyInfo
	_, err = asn1.Unmarshal(publicKey, &subPKI)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %s", err)
	}

	h := sha1.New()
	h.Write(subPKI.SubjectPublicKey.Bytes)
	return h.Sum(nil), nil
}

func savePrivateKey(key *rsa.PrivateKey, path string, password []byte) error {
	b := x509.MarshalPKCS1PrivateKey(key)
	var block *pem.Block
	var err error

	if len(password) > 3 {
		block, err = x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", b, password, x509.PEMCipherAES256)
		if err != nil {
			return fmt.Errorf("Unable to encrypt key: %s", err)
		}
	} else {
		block = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: b}
	}

	keyOut, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open key.pem for writing: %v", err)
	}

	pem.Encode(keyOut, block)
	keyOut.Close()

	return nil
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
