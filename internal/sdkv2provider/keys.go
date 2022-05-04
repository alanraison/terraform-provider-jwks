package sdkv2provider

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"golang.org/x/crypto/ssh"
)

func readKeyData(data string) (interface{}, error) {
	var keyData interface{}
	dataBytes := []byte(data)
	b64data, err := base64.StdEncoding.DecodeString(data)
	if err == nil {
		dataBytes = b64data
	}
	block, _ := pem.Decode(dataBytes)
	if block != nil {
		//handle pem encoded
		keyData, err = ssh.ParseRawPrivateKey(dataBytes)
		if err != nil {
			keyData, err = x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return nil, errors.New("unable to parse private or public key pem")
			}
		}
	} else {
		keyData, err = x509.ParsePKCS8PrivateKey(dataBytes)
		if err != nil {
			keyData, err = x509.ParsePKCS1PrivateKey(dataBytes)
			if err != nil {
				keyData, err = x509.ParseECPrivateKey(dataBytes)
				if err != nil {
					keyData, err = x509.ParsePKIXPublicKey(dataBytes)
					if err != nil {
						return nil, errors.New("unable to parse private or public key pem")
					}
				}
			}
		}
	}
	return keyData, nil
}
