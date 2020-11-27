package jw_token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"go_training/config"
	"os"
)

// 参考
// https://gist.github.com/sdorra/1c95de8cb80da31610d2ad767cd6f251
func KeyGenerator(conf config.Config) {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	mustKeyGen(err)

	publicKey := key.PublicKey
	keyPath := conf.App.KeyPath
	saveGobKey(keyPath+"private.key", key)
	savePEMKey(keyPath+"private.pem", key)

	saveGobKey(keyPath+"public.key", publicKey)
	savePublicPEMKey(keyPath+"public.pem", publicKey)
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	mustKeyGen(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	mustKeyGen(err)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	mustKeyGen(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	mustKeyGen(err)
}

func savePublicPEMKey(fileName string, pubKey rsa.PublicKey) {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(&pubKey)
	mustKeyGen(err)

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemFile, err := os.Create(fileName)
	mustKeyGen(err)
	defer pemFile.Close()

	err = pem.Encode(pemFile, pemKey)
	mustKeyGen(err)
}

func mustKeyGen(err error) {
	if err != nil {
		panic(fmt.Sprintf("key generate error: %s", err.Error()))
	}
}
