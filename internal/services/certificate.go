// internal/services/certificate.go
package services

import (
	"crypto/x509"
	"os"

	"software.sslmate.com/src/go-pkcs12"
)

type Certificate struct {
	Cert    *x509.Certificate
	PrivKey interface{}
}

func LoadCertificate(certPath, password string) (*Certificate, error) {
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	// Parse PKCS12 format
	privateKey, cert, err := pkcs12.Decode(certData, password)
	if err != nil {
		return nil, err
	}

	return &Certificate{
		Cert:    cert,
		PrivKey: privateKey,
	}, nil
}
