package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

const (
	privateKeyFile = "private.pem"
	publicKeyFile  = "public.pem"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func Init() (*rsa.PublicKey, error) {
	var err error

	if _, err = os.Stat(privateKeyFile); err != nil {
		privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, err
		}

		privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privKeyPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		})
		err = os.WriteFile(privateKeyFile, privKeyPem, 0600)
		if err != nil {
			return nil, err
		}
	} else {
		privKeyPem, err := os.ReadFile(privateKeyFile)
		if err != nil {
			return nil, err
		}
		privKeyBlock, _ := pem.Decode(privKeyPem)
		privateKey, err = x509.ParsePKCS1PrivateKey(privKeyBlock.Bytes)
		if err != nil {
			return nil, err
		}
	}

	if _, err = os.Stat(publicKeyFile); err != nil {
		publicKey = &privateKey.PublicKey

		pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			return nil, err
		}
		pubKeyPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKeyBytes,
		})
		err = os.WriteFile(publicKeyFile, pubKeyPem, 0644)
		if err != nil {
			return nil, err
		}
	} else {
		pubKeyPem, err := os.ReadFile(publicKeyFile)
		if err != nil {
			return nil, err
		}
		pubKeyBlock, _ := pem.Decode(pubKeyPem)
		pubKeyIface, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
		if err != nil {
			return nil, err
		}
		publicKey = pubKeyIface.(*rsa.PublicKey)
	}

	return publicKey, nil
}

func Encrypt(data []byte) ([]byte, error) {
	if publicKey == nil {
		return nil, errors.New("encryption public key not found")
	}
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

func Decrypt(data []byte) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("decryption private key not found")
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
}

func PublicKeyToString(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return string(publicKeyPem), nil
}

func StringToPublicKey(publicKeyString string) (*rsa.PublicKey, error) {
	publicKeyPem, _ := pem.Decode([]byte(publicKeyString))
	if publicKeyPem == nil {
		return nil, errors.New("failed to decode public key PEM")
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyPem.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("parsed public key is not an RSA public key")
	}

	return rsaPublicKey, nil
}
