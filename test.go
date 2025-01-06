package mpesasdk
//
//import (
//	"bytes"
//	"encoding/json"
//	"io"
//	"net/http"
//	"github.com/coleYab/mpesasdk/utils"
//	"github.com/coleYab/mpesasdk/auth"
//)
//
//const (
//	authTypeBearer = "Bearer"
//	authTypeBasic  = "Basic"
//)
//
//type HttpClient struct {
//	client *http.Client
//	auth   *auth.AuthorizationToken // dependency on AuthorizationToken for Bearer token
//}
//
//func NewHttpClient(timeout uint, maxRetries uint, auth *auth.AuthorizationToken) *HttpClient {
//	if timeout == 0 {
//		timeout = 1
//	}
//
//	return &HttpClient{
//		client: &http.Client{
//			Timeout: time.Duration(timeout) * time.Second,
//		},
//		auth: auth, // inject the authorization token here
//	}
//}
//
//func (c *HttpClient) ApiRequest(env, endpoint, method string, payload interface{}, authType string) (*http.Response, error) {
//	url := utils.ConstructURL(env, endpoint)
//
//	var body io.Reader
//	if payload != nil {
//		jsonData, err := json.Marshal(payload)
//		if err != nil {
//			return nil, err
//		}
//		body = bytes.NewReader(jsonData)
//	}
//
//	req, err := http.NewRequest(method, url, body)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Add("Content-Type", "application/json")
//
//	// Handle authorization type
//	switch authType {
//	case authTypeBearer:
//		// Use the AuthorizationToken to fetch the Bearer token
//		authToken, err := c.auth.GetAuthorizationToken(string(env), c.auth.ConsumerKey, c.auth.ConsumerSecret)
//		if err != nil {
//			return nil, err
//		}
//		req.Header.Add("Authorization", authToken)
//	case authTypeBasic:
//		req.SetBasicAuth(c.auth.ConsumerKey, c.auth.ConsumerSecret)
//	}
//
//	// Perform the API request
//	res, err := c.client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	return res, nil
//}
//
////
//package mpesasdk
//
//import (
//	"errors"
//	"github.com/coleYab/mpesasdk/auth"
//	"github.com/coleYab/mpesasdk/client"
//	"github.com/coleYab/mpesasdk/service"
//	"github.com/coleYab/mpesasdk/utils"
//	"time"
//)
//
//type Environment string
//
//const (
//	PRODUCTION Environment = "Production"
//	SANDBOX    Environment = "SandBox"
//)
//
//type MpesaClient struct {
//	consumerKey    string
//	consumerSecret string
//	authorizationToken *auth.AuthorizationToken
//	env            Environment
//	client         *client.HttpClient
//	logger         *service.Logger
//}
//
//func NewMpesaClient(consumerKey, consumerSecret string, env Environment, logLevel service.LogLevel, timeout time.Duration, maxRetries uint) (*MpesaClient, error) {
//	if env != PRODUCTION && env != SANDBOX {
//		return nil, errors.New("invalid environment: must be either 'Production' or 'Sandbox'")
//	}
//
//	if consumerKey == "" || consumerSecret == "" {
//		return nil, errors.New("consumer key and consumer secret cannot be empty")
//	}
//
//	if timeout <= 0 {
//		timeout = 5 * time.Second
//	}
//
//	if maxRetries < 0 {
//		maxRetries = 1
//	}
//
//	// Initialize AuthorizationToken and HttpClient
//	authToken := &auth.AuthorizationToken{}
//	httpClient := client.NewHttpClient(uint(timeout.Seconds()), maxRetries, authToken)
//
//	// Create and return the MpesaClient
//	logger := service.NewLogger(logLevel)
//
//	return &MpesaClient{
//		consumerKey:       consumerKey,
//		consumerSecret:    consumerSecret,
//		authorizationToken: authToken,
//		env:               env,
//		client:            httpClient,
//		logger:            logger,
//	}, nil
//}
//
