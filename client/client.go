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

type HttpClient struct {
	client *http.Client
	auth   *auth.AuthorizationToken // dependency on AuthorizationToken for Bearer and Basic token
    maxRetries uint
}

func NewHttpClient(timeout time.Duration, maxRetries uint, auth *auth.AuthorizationToken) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
        maxRetries: maxRetries,
		auth: auth,
	}
}


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

	for attempt := uint(0); attempt <= c.maxRetries; attempt++ {
		res, err = c.makeRequest(url, method, body, authType, env)
		if err == nil || !isTimeoutError(err) {
			// If no error or it's not a timeout error, return the response or error
			break
		}

		// Add a delay between retries
		time.Sleep(time.Duration(attempt+1) * time.Second)
	}

	return res, err
}

func (c *HttpClient) makeRequest(url, method string, body io.Reader, authType string, env common.Enviroment) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	// Handle authorization type
	switch authType {
	case auth.AuthTypeBearer:
		// Use the AuthorizationToken to fetch the Bearer token
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


// isTimeoutError checks if the error is related to a timeout.
func isTimeoutError(err error) bool {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	return errors.Is(err, context.DeadlineExceeded)
}
