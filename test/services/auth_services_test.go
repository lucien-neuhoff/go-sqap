package services_test

import (
	"bytes"
	"encoding/json"
	"go-sqap/internal/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// Initialize the test server and client
	router := gin.Default()
	server := httptest.NewServer(router)
	defer server.Close()

	client := server.Client()

	// Create the request payload
	createReq := models.CreateUserRequest{
		Email:    "john.doe@example.com",
		Password: "Password123",
	}

	// Marshal the request payload into JSON
	reqBody, err := json.Marshal(createReq)
	require.NoError(t, err)
	// Create the HTTP request
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/api/register", bytes.NewBuffer(reqBody))
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

	// Unmarshal the response body into a User struct
	var user models.User
	err = json.Unmarshal(respBody, &user)
	require.NoError(t, err)

	// Assert that the user's fields match the request payload
	require.Equal(t, createReq.Email, user.Email)
}

func TestLoginUser(t *testing.T) {
	// Initialize the test server and client
	router := gin.Default()
	server := httptest.NewServer(router)
	defer server.Close()

	client := server.Client()

	// Create the request payload
	createReq := models.CreateUserRequest{
		Email:    "john.doe@example.com",
		Password: "Password123",
	}

	// Marshal the request payload into JSON
	reqBody, err := json.Marshal(createReq)
	require.NoError(t, err)
	// Create the HTTP request
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/api/login", bytes.NewBuffer(reqBody))
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
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Unmarshal the response body into a User struct
	var response struct {
		User struct {
			UUID  string `json:"uuid"`
			Email string `json:"email"`
		} `json:"user"`
		Token string `json:"token"`
	}
	err = json.Unmarshal(respBody, &response)
	require.NoError(t, err)

	// Assert that the user's fields match the request payload
	require.Equal(t, createReq.Email, response.User.Email)
}
