package services_test

import (
	"crypto/rand"
	"crypto/rsa"
	"go-sqap/encryption"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryption(t *testing.T) {
	sPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	require.NoError(t, err)
	sPublicKey := sPrivateKey.PublicKey

	cPrivateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	require.NoError(t, err)
	cPublicKey := cPrivateKey.PublicKey

	sMessage := "hello"

	esMessage, err := encryption.Encrypt([]byte(sMessage), &sPublicKey)
	require.NoError(t, err)

	dsMessage, err := encryption.Decrypt([]byte(esMessage), sPrivateKey)
	require.NoError(t, err)

	require.Equal(t, string(dsMessage), sMessage)

	cMessage := "world"

	ecMessage, err := encryption.Encrypt([]byte(cMessage), &cPublicKey)
	require.NoError(t, err)

	deMessage, err := encryption.Decrypt([]byte(ecMessage), cPrivateKey)
	require.NoError(t, err)

	require.Equal(t, string(deMessage), cMessage)
}
