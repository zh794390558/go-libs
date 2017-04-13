package main

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func GenRSAPriv(fileName, passwd string, len int) error {
	priv, err := rsa.GenerateKey(rand.Reader, len)
	if err != nil {
		return err

	}
	fmt.Println("RSA Private Key:", priv)

	//converts a private key to ASN.1 DER encoded form.
	data := x509.MarshalPKCS1PrivateKey(priv)
	fmt.Println("ASN.1 DER:", data)
	err = encodePrivPemFile(fileName, passwd, data)
	return err

}

//GenECDSAPriv 生成ECDSA私钥文件
func GenECDSAPriv(fileName, passwd string) error {
	priv, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		return err

	}
	data, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return err

	}
	err = encodePrivPemFile(fileName, passwd, data)
	return err

}

//GenDSAPriv 生成DSA私钥(用于演示)
func GenDSAPriv() {
	priv := &dsa.PrivateKey{}
	dsa.GenerateParameters(&priv.Parameters, rand.Reader, dsa.L1024N160)
	dsa.GenerateKey(priv, rand.Reader)
	fmt.Printf("priv:%+vn", priv)

}

//DecodePriv 解析私钥文件生成私钥，（RSA，和ECDSA两种私钥格式）
func DecodePriv(fileName, passwd string) (pubkey, priv interface{}, err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, nil, errors.New("读取私钥文件错误")

	}
	block, _ := pem.Decode(data)
	data, err = x509.DecryptPEMBlock(block, []byte(passwd))
	if err != nil {
		return nil, nil, err

	}

	privKey, err := x509.ParsePKCS1PrivateKey(data) //解析成RSA私钥
	if err != nil {
		priv, err = x509.ParseECPrivateKey(data) //解析成ECDSA私钥
		if err != nil {
			return nil, nil, errors.New("支持持RSA和ECDSA格式的私钥")

		}

	}
	priv = privKey
	pubkey = &privKey.PublicKey
	return

}

//生成私钥的pem文件
func encodePrivPemFile(fileName, passwd string, data []byte) error {
	//encrypt EDR-encoded file with passwd under cipher, out form is PEM
	block, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", data, []byte(passwd), x509.PEMCipher3DES)
	if err != nil {
		return err

	}
	fmt.Println("PEM Block:", block)
	file, err := os.Create(fileName)
	if err != nil {
		return err

	}
	err = pem.Encode(file, block)
	if err != nil {
		return err

	}
	return nil

}

// EncodeCsr 生成证书请求
func EncodeCsr(country, organization, organizationlUnit, locality, province, streetAddress, postallCode []string, commonName, fileName string, priv interface{}) error {
	req := &x509.CertificateRequest{
		Subject: pkix.Name{
			Country:            country,
			Organization:       organization,
			OrganizationalUnit: organizationlUnit,
			Locality:           locality,
			Province:           province,
			StreetAddress:      streetAddress,
			PostalCode:         postallCode,
			CommonName:         commonName,
		},
	}

	data, err := x509.CreateCertificateRequest(rand.Reader, req, priv)
	if err != nil {
		return err

	}
	err = util.EncodePemFile(fileName, "CERTIFICATE REQUEST", data)
	return err

}

//DecodeCsr 解析CSRpem文件
func DecodeCsr(fileName string) (*x509.CertificateRequest, error) {
	data, err := util.DecodePemFile(fileName)
	if err != nil {
		return nil, err

	}

	req, err := x509.ParseCertificateRequest(data)
	return req, err

}

//GenSignselfCertificate 生成自签名证书
func GenSignselfCertificate(req *x509.CertificateRequest, publickey, privKey interface{}, fileName string, maxPath int, days time.Duration) error {
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(random.Int63n(time.Now().Unix())),
		Subject:               req.Subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(days * 24 * time.Hour),
		BasicConstraintsValid: true,
		IsCA:               true,
		SignatureAlgorithm: x509.SHA1WithRSA, // 签名算法选择SHA1WithRSA
		KeyUsage:           x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDataEncipherment,
		SubjectKeyId:       []byte{1, 2, 3},
	}
	if maxPath > 0 { //如果长度超过0则设置了 最大的路径长度
		template.MaxPathLen = maxPath
	}
	cert, err := x509.CreateCertificate(rand.Reader, template, template, publickey, privKey)
	if err != nil {
		return errors.New("签发自签名证书失败")

	}
	err = util.EncodePemFile(fileName, "CERTIFICATE", cert)
	if err != nil {
		return err

	}
	return nil

}

//GenCertificate 生成非自签名证书
func GenCertificate(req *x509.CertificateRequest, parentCert *x509.Certificate, pubKey, parentPrivKey interface{}, fileName string, isCA bool, days time.Duration) error {
	template := &x509.Certificate{
		SerialNumber: big.NewInt(random.Int63n(time.Now().Unix())),
		Subject:      req.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(days * 24 * time.Hour),
		// ExtKeyUsage: []x509.ExtKeyUsage{ //额外的使用
		//  x509.ExtKeyUsageClientAuth,
		//  x509.ExtKeyUsageServerAuth,
		// },
		//

		SignatureAlgorithm: x509.SHA1WithRSA,
	}

	if isCA {
		template.BasicConstraintsValid = true
		template.IsCA = true

	}

	cert, err := x509.CreateCertificate(rand.Reader, template, parentCert, pubKey, parentPrivKey)
	if err != nil {
		return errors.New("签署证书失败")

	}
	err = util.EncodePemFile(fileName, "CERTIFICATE", cert)
	if err != nil {
		return err

	}
	return nil

}

func main() {
	GenRSAPriv("rsa.key", "abc", 2048)
}
