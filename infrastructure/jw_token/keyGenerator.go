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

type KeyGenerator struct {
	path       string
	privateKey *rsa.PrivateKey
	pubKey     *rsa.PublicKey
}

func NewKeyGenerator(path string) KeyGenerator {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	mustKeyGen(err)

	return KeyGenerator{path: path, privateKey: key, pubKey: &key.PublicKey}
}

func (g KeyGenerator) Generate(conf config.Config) {
	g.savePrivateKey()
	g.savePEMKey()
	g.savePublicKey()
	g.savePublicPEMKey()
}

func (g KeyGenerator) savePrivateKey() {
	outFile, err := os.Create(g.path + "private.key")
	mustKeyGen(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(g.privateKey)
	mustKeyGen(err)
}

func (g KeyGenerator) savePublicKey() {
	outFile, err := os.Create(g.path + "public.key")
	mustKeyGen(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(g.pubKey)
	mustKeyGen(err)
}

func (g KeyGenerator) savePEMKey() {
	outFile, err := os.Create(g.path + "private.pem")
	mustKeyGen(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(g.privateKey),
	}

	err = pem.Encode(outFile, privateKey)
	mustKeyGen(err)
}

func (g KeyGenerator) savePublicPEMKey() {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(g.pubKey)
	mustKeyGen(err)

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemFile, err := os.Create(g.path + "public.pem")
	mustKeyGen(err)
	defer pemFile.Close()

	err = pem.Encode(pemFile, pemKey)
	mustKeyGen(err)
}

func mustKeyGen(err error) {
	if err != nil {
		panic(fmt.Sprintf("privateKey generate error: %s", err.Error()))
	}
}
