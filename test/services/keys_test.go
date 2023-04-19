package services_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"go-sqap/encryption"
	"go-sqap/internal/models"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	// Initialize the test server and client
	router := gin.Default()
	server := httptest.NewServer(router)
	defer server.Close()

	client := server.Client()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error while generating private key: ", err)
		return
	}

	publicKeyString, err := encryption.PublicKeyToString(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("Error while parsing publicKey to string: ", err)
		return
	}

	// Create the request payload
	createReq := models.PublicKeyRequest{
		Email: "john.doe@example.com",
		Key:   publicKeyString,
	}

	// Marshal the request payload into JSON
	reqBody, err := json.Marshal(createReq)
	require.NoError(t, err)
	// Create the HTTP request
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/api/keys", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	// Assert that the response status code is 201 Created
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	// Unmarshal the response body into a map
	var response map[string]string
	err = json.Unmarshal(respBody, &response)
	require.NoError(t, err)

	// Assert that the server's public key is returned
	serverPublicKeyStr, ok := response["server_public_key"]
	require.True(t, ok)

	serverPublicKey, err := encryption.StringToPublicKey(serverPublicKeyStr)
	require.NoError(t, err)

	keyData, err := os.ReadFile("../../public.pem")
	require.NoError(t, err)

	// Parse the PEM block containing the public key
	block, _ := pem.Decode(keyData)

	// Parse the RSA public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	require.NoError(t, err)

	require.Equal(t, publicKey, serverPublicKey)
}
