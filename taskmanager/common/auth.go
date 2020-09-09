package common

import (
	"io/ioutil"
)

// using asymetric crypto/RSA key
const (
	PRIVATE_KEY_PATH = "keys/app.rsa"
	PUBLIC_KEY_PATH = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// Read the key files before starting http handlers
func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(PRIVATE_KEY_PATH)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	verifyKey, err = ioutil.ReadFile(PUBLIC_KEY_PATH)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}
}