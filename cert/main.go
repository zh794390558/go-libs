package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
)

var (
	crt_f   = flag.String("crt", "", "user crt file")
	key_f   = flag.String("key", "", "user rkey file")
	cacrt_f = flag.String("ca-crt", "", "ca crt file")
	cakey_f = flag.String("ca-key", "", "ca key file")
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
func ParsePKCS1PrivateKey(filename string) (*rsa.PrivateKey, error) {
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

func publicKey(priv interface{}) interface{} {
	switch priv.(type) {
	case *rsa.PrivateKey:
		return &priv.(*rsa.PrivateKey).PublicKey
	case *ecdsa.PrivateKey:
		return &priv.(*ecdsa.PrivateKey).PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
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

//create certificate for user withou self-sign
func CreateUserCertificate(caCrt *x509.Certificate, cakey interface{}, key interface{}, userName string, t time.Duration) (cert []byte) {
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

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, caCrt, publicKey(key), cakey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	return derBytes
}

func main() {
	flag.Parse()

	crt, _ := ParseCertificate(*crt_f)

	key, _ := ParsePKCS1PrivateKey(*key_f)
	key = key

	cacrt, _ := ParseCertificate(*cacrt_f)

	cakey, err := ParsePKCS1PrivateKey(*cakey_f)
	if err != nil {
		log.Fatal(err)
	}

	if err := crt.CheckSignatureFrom(cacrt); err != nil {
		fmt.Println("ok")
	} else {
		fmt.Println("err")
	}

	var priv interface{}
	priv, err = rsa.GenerateKey(rand.Reader, 2048)

	derBytes := CreateUserCertificate(cacrt, cakey, priv, "testUser", 24*time.Hour)

	certOut, _ := os.Create("cert.pem")
	defer certOut.Close()
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut, err := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Print("failed to open key.pem for writing:", err)
		return

	}
	pem.Encode(keyOut, pemBlockForKey(priv))
	keyOut.Close()

	return
}
