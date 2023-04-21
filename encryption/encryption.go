package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"fmt"
	"go-sqap/internal/models"
	"os"
)

const (
	privateKeyFile = "private.pem"
	publicKeyFile  = "public.pem"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func Init() error {
	var err error

	if _, err = os.Stat(privateKeyFile); err != nil {
		privateKey, err = rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			return err
		}

		privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privKeyPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		})
		f, err := os.Create(privateKeyFile)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(privKeyPem)
		if err != nil {
			return err
		}
	} else {
		privKeyPem, err := os.ReadFile(privateKeyFile)
		if err != nil {
			return err
		}
		privKeyBlock, _ := pem.Decode(privKeyPem)
		privateKey, err = x509.ParsePKCS1PrivateKey(privKeyBlock.Bytes)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(publicKeyFile); err != nil {
		publicKey = &privateKey.PublicKey

		pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			return err
		}
		pubKeyPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKeyBytes,
		})

		f, err := os.Create(publicKeyFile)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(pubKeyPem)
		if err != nil {
			return err
		}
	} else {
		pubKeyPem, err := os.ReadFile(publicKeyFile)
		if err != nil {
			return err
		}
		pubKeyBlock, _ := pem.Decode(pubKeyPem)
		pubKeyIface, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
		if err != nil {
			return err
		}
		publicKey = pubKeyIface.(*rsa.PublicKey)
	}

	return nil
}

func Encrypt(data []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

func EncryptS(data []byte, publicKeyString string) ([]byte, error) {
	userPublicKey, err := StringToPublicKey(publicKeyString)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, userPublicKey, data)
}

func EncryptUser(user models.User, publicKey *rsa.PublicKey) (*models.EncryptedUser, error) {
	uuidBytes := []byte(user.UUID)
	emailBytes := []byte(user.Email)

	createdAtBytes := make([]byte, user.CreatedAt.Time.Unix())
	binary.BigEndian.PutUint64(createdAtBytes, uint64(user.CreatedAt.Time.Unix()))

	encryptedUUID, err := Encrypt(uuidBytes, publicKey)
	if err != nil {
		return nil, err
	}

	encryptedEmail, err := Encrypt(emailBytes, publicKey)
	if err != nil {
		return nil, err
	}

	encryptedCreatedAt, err := Encrypt(createdAtBytes, publicKey)
	if err != nil {
		return nil, err
	}

	encryptedUser := models.EncryptedUser{UUID: encryptedUUID, Email: encryptedEmail, CreatedAt: encryptedCreatedAt}
	return &encryptedUser, nil
}

func Decrypt(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("decryption private key not found")
	}
	fmt.Println("Decrypting\n", data, "\nWith\n", privateKey)
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

func GetServerPublicKey() *rsa.PublicKey {
	return publicKey
}

func GetServerPublicKeyString() (*string, error) {
	publicKeyStr, err := PublicKeyToString(publicKey)
	if err != nil {
		return nil, err
	}

	return &publicKeyStr, nil
}
