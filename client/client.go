package client

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
	"net/http"
	"github.com/coleYab/mpesasdk/utils"
	"github.com/coleYab/mpesasdk/auth"
)

type HttpClient struct {
	client *http.Client
	auth   *auth.AuthorizationToken // dependency on AuthorizationToken for Bearer and Basic token
}

func NewHttpClient(timeout time.Duration, maxRetries uint, auth *auth.AuthorizationToken) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
		auth: auth,
	}
}

func (c *HttpClient) ApiRequest(env, endpoint, method string, payload interface{}, authType string) (*http.Response, error) {
	url := utils.ConstructURL(env, endpoint)

	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonData)
	}

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
		authToken, err := c.auth.GetAuthorizationToken(string(env), key, secret)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", authToken)
	case auth.AuthTypeBasic:
		req.SetBasicAuth(c.auth.GetConsumerKeyAndSecret())
	}

	// Perform the API request
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
