package main

import (
	"flag"
	"fmt"
	log "github.com/golang/glog"
	"github.com/zh794390558/go-study/cert/gencrt"
	"time"
)

var (
	crt_f   = flag.String("crt", "", "user crt file")
	key_f   = flag.String("key", "", "user rkey file")
	cacrt_f = flag.String("ca-crt", "", "ca crt file")
	cakey_f = flag.String("ca-key", "", "ca key file")
)

func main() {
	flag.Parse()
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
	gencrt.PemEncode("key.pem", gencrt.PemBlockForKey(priv))

	//create certificate withou self-sign
	derBytes := gencrt.CreateUserCertificate(cacrt, cakey, priv, "testUser", 24*time.Hour)
	gencrt.PemEncode("cert.pem", gencrt.PemBlockForCrt(derBytes))

	return
}
