// Package client provides an HTTP client implementation for interacting with the M-Pesa API.
// It handles API requests, retries on failure, and authorization using Bearer and Basic tokens.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/coleYab/mpesasdk/auth"
	"github.com/coleYab/mpesasdk/common"
	"github.com/coleYab/mpesasdk/utils"
)

// HttpClient is a custom HTTP client designed to interact with the M-Pesa API.
// It supports request retries, timeout handling, and token-based authentication.
//
// Fields:
//   - client: The underlying http.Client instance used for making requests.
//   - auth: An instance of AuthorizationToken used to handle authentication.
//   - maxRetries: The maximum number of retry attempts for timeout errors.
type HttpClient struct {
	client     *http.Client
	auth       *auth.AuthorizationToken
	maxRetries uint
}

// NewHttpClient creates a new instance of HttpClient.
//
// Parameters:
//   - timeout: The timeout duration for each HTTP request.
//   - maxRetries: The maximum number of retry attempts for timeout errors.
//   - auth: An instance of AuthorizationToken for managing authentication.
//
// Returns:
//   - A pointer to the initialized HttpClient.
func NewHttpClient(timeout time.Duration, maxRetries uint, auth *auth.AuthorizationToken) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
		maxRetries: maxRetries,
		auth:       auth,
	}
}

// ApiRequest sends an HTTP request to the specified M-Pesa API endpoint.
//
// Parameters:
//   - env: The environment (sandbox or production) to determine the base URL.
//   - endpoint: The API endpoint to call.
//   - method: The HTTP method (e.g., "GET", "POST").
//   - payload: The request payload, serialized to JSON.
//   - authType: The type of authorization to use (e.g., "Bearer", "Basic").
//
// Returns:
//   - *http.Response: The HTTP response from the server.
//   - error: Any error encountered during the request.
func (c *HttpClient) ApiRequest(env common.Enviroment, endpoint, method string, payload interface{}, authType string) (*http.Response, error) {
	url := utils.ConstructURL(env, endpoint)

	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonData)
	}

	var res *http.Response
	var err error

	// Retry loop for handling timeout errors
	for attempt := uint(0); attempt <= c.maxRetries; attempt++ {
		res, err = c.makeRequest(url, method, body, authType, env)
		if err == nil || !isTimeoutError(err) {
			break
		}

		// Add a delay before the next retry
		time.Sleep(time.Duration(attempt+1) * time.Second)
	}

	return res, err
}

// makeRequest constructs and sends an HTTP request with the given parameters.
//
// Parameters:
//   - url: The full URL of the API endpoint.
//   - method: The HTTP method (e.g., "GET", "POST").
//   - body: The request body, if applicable.
//   - authType: The type of authorization to use (e.g., "Bearer", "Basic").
//   - env: The environment (sandbox or production).
//
// Returns:
//   - *http.Response: The HTTP response from the server.
//   - error: Any error encountered during the request.
func (c *HttpClient) makeRequest(url, method string, body io.Reader, authType string, env common.Enviroment) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	// Handle authorization based on the specified authType
	switch authType {
	case auth.AuthTypeBearer:
		key, secret := c.auth.GetConsumerKeyAndSecret()
		authToken, err := c.auth.GetAuthorizationToken(env, key, secret)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", authToken)
	case auth.AuthTypeBasic:
		req.SetBasicAuth(c.auth.GetConsumerKeyAndSecret())
	}

	return c.client.Do(req)
}

// isTimeoutError checks whether the given error is related to a timeout.
//
// Parameters:
//   - err: The error to check.
//
// Returns:
//   - bool: True if the error is a timeout error, false otherwise.
func isTimeoutError(err error) bool {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return errors.Is(err, context.DeadlineExceeded)
}

