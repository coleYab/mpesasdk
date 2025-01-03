package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
    authTypeBearer = "Bearer"
    authTypeBasic  = "Basic"
)

type HttpClient struct {
	client *http.Client
    maxRetries uint
};

func NewHttpClient(timeout, maxRetries uint) *HttpClient {
    if timeout == 0 {
        timeout = 1;
    }

    return &HttpClient{
        client: &http.Client{
            Timeout: time.Duration(timeout) * time.Second,
        },

        maxRetries: maxRetries,
    }
}

func (c *HttpClient) MakeRequest(url, method string, payload interface{}, authType string) ([]byte, error) {
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
    switch authType {
    case authTypeBearer:
        authToken, err := getAuthorizationToken()
        if err != nil {
            return nil, err
        }
        req.Header.Add("Authorization", authToken)
    case authTypeBasic:
        req.SetBasicAuth(client.consumerKey, client.consumerSecret)
    }

    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    responseData, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    return responseData, nil
}


