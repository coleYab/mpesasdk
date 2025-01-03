package client

import (
	"net/http"
	"strings"
	"time"
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

func (c *HttpClient) MakeRequest(url, method, payload string, authToken string) (*http.Response, error) {
    req, err := http.NewRequest(method, url, strings.NewReader(payload))
    if err != nil {
        return nil, err
    }

    req.Header.Add("Content-Type", "application/json")

    if authToken != "" {
        req.Header.Add("Authorization", authToken)
    }

    res, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    return res, nil
}


