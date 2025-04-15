// internal/services/signer.go
package services

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/beevik/etree"
)

type Signer struct {
	cert *Certificate
}

func NewSigner(cert *Certificate) *Signer {
	return &Signer{cert: cert}
}

func (s *Signer) SignXML(xmlDoc *etree.Document) error {
	// Find the Signature element
	signature := xmlDoc.FindElement("//Signature")
	if signature == nil {
		return fmt.Errorf("signature element not found")
	}

	// Calculate digest value
	signedInfo := signature.FindElement("SignedInfo")
	if signedInfo == nil {
		return fmt.Errorf("SignedInfo element not found")
	}

	// Canonicalize SignedInfo
	canonicalSignedInfo := signedInfo.Text()

	// Calculate signature
	hash := sha1.New()
	hash.Write([]byte(canonicalSignedInfo))
	digest := hash.Sum(nil)

	privateKey, ok := s.cert.PrivKey.(*rsa.PrivateKey)
	if !ok {
		return fmt.Errorf("private key is not RSA")
	}

	signature_value, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, digest)
	if err != nil {
		return err
	}

	// Add signature value to XML
	signatureValue := signature.FindElement("SignatureValue")
	if signatureValue == nil {
		return fmt.Errorf("SignatureValue element not found")
	}
	signatureValue.SetText(base64.StdEncoding.EncodeToString(signature_value))

	return nil
}
