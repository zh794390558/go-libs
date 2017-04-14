package gencrt

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
)

var (
	ErrorPemDecode = errors.New("Pem decode error")
)

//candy x509.ParseCertificate func
func ParseCertificate(filename string) (*x509.Certificate, error) {
	blk, err := PemDecode(filename)
	if err != nil {
		return nil, err
	}

	crt, err := x509.ParseCertificate(blk.Bytes)
	if err != nil {
		return nil, err
	}

	return crt, nil
}

//candy x509.Parsepkcs1privatekey func
func ParseRSAPrivateKey(filename string) (*rsa.PrivateKey, error) {
	blk, err := PemDecode(filename)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(blk.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

//candy pem.Decode func
func PemDecode(filename string) (*pem.Block, error) {
	var data []byte
	var err error
	var blk *pem.Block

	if data, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}

	if blk, data = pem.Decode(data); blk == nil {
		return nil, ErrorPemDecode
	}

	return blk, nil
}

//candy pem.Encode func
func PemEncode(filename string, blk *pem.Block) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = pem.Encode(f, blk)
	if err != nil {
		return err
	}
	return nil
}

func PemEncodeToMemory(blk *pem.Block) []byte {
	return pem.EncodeToMemory(blk)
}

//get public key
func PublicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// ANS.1 DER-encode
func DerForKey(priv interface{}) ([]byte, error) {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return x509.MarshalPKCS1PrivateKey(k), nil
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, err
		}
		return b, nil
	default:
		return nil, errors.New("Do not support this type priv key")
	}
}

func PemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		fmt.Println(k)
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)

		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

func PemBlockForCrt(crtBytes []byte) *pem.Block {
	return &pem.Block{Type: "CERTIFICATE", Bytes: crtBytes}
}

//create certificate for user withou self-sign
func CreateUserCertificate(caCrt *x509.Certificate, cakey interface{}, userKey interface{}, userName string, t time.Duration) (cert []byte) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)

	}

	notBefore := time.Now()
	notAfter := notBefore.Add(t)

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: userName,
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: false,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, caCrt, PublicKey(userKey), cakey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	return derBytes
}

func GenerateRSAPrivKey(bits int) (*rsa.PrivateKey, error) {
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return priv, nil
}
