package main

import (
	"fmt"
	"github.com/zh794390558/go-study/cert/gencrt"
	"log"
	"time"
)

func main() {
	//read user crt
	crt, _ := gencrt.ParseCertificate(*crt_f)

	//read user key
	key, _ := gencrt.ParseRSAPrivateKey(*key_f)
	key = key

	//read ca crt
	cacrt, _ := gencrt.ParseCertificate(*cacrt_f)

	//read ca key
	cakey, err := gencrt.ParseRSAPrivateKey(*cakey_f)
	if err != nil {
		log.Fatal(err)
	}

	if err := crt.CheckSignatureFrom(cacrt); err != nil {
		fmt.Println("ok")
	} else {
		fmt.Println("err")
	}

	//create priv key
	var priv interface{}
	priv, err = gencrt.GenerateRSAPrivKey(2048)
	gencrt.PemEncode("key.pem", PemBlockForKey(priv))

	//create certificate withou self-sign
	derBytes := gencrt.CreateUserCertificate(cacrt, cakey, priv, "testUser", 24*time.Hour)
	gencrt.PemEncode("cert.pem", PemBlockForCrt(derBytes))

	return
}
