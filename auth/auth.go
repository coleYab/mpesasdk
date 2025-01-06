package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"github.com/coleYab/mpesasdk/utils"
)

type AuthorizationToken struct {
	token     string // Bearer
	createdAt time.Time
	expiresIn int
    consumerKey string
    consumerSecret string
}

const (
    AuthTypeBearer = "Bearer"
    AuthTypeNone = ""
    AuthTypeBasic  = "Basic"
)

func NewAuthorizationToken(key, secret string) *AuthorizationToken {
    return &AuthorizationToken{
        consumerKey: key,
        consumerSecret: secret,
    }
}


func (a *AuthorizationToken) setAuthToken(tokenType, token string, expiresIn int, key, secret string) {
	a.token = fmt.Sprintf("%v %v", tokenType, token)
	a.createdAt = time.Now()
	a.expiresIn = expiresIn
    a.consumerKey = key
    a.consumerSecret = secret
}

func (a *AuthorizationToken) GetConsumerKeyAndSecret() (string, string) {
    return a.consumerKey, a.consumerSecret
}

func (a *AuthorizationToken) GetAuthorizationToken(env string, key, secret string) (string, error) {
	url := utils.ConstructURL(env, "/v1/token/generate?grant_type=client_credentials")
	method := "GET"

	// If token is still valid (2 seconds before expiry), return it
	if a.token != "" && time.Now().Before(a.createdAt.Add(time.Duration(a.expiresIn-2)*time.Second)) {
		return a.token, nil
	}

	// Otherwise, request a new token
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", errors.New("error: while creating auth request")
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(key, secret)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var authResponse struct {
        AccessToken string `json:"access_token"`
        TokenType   string `json:"token_type"`
        ExpiresIn   string `json:"expires_in"`
        ResultCode  string `json:"resultCode"`
        ResultDesc  string `json:"resultDesc"`
    }

	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		return "", err
	}

	// Handle errors from the response
	if authResponse.ResultCode != "" {
		return "", errors.New(authResponse.ResultDesc)
	}

	expiresIn, _ := strconv.Atoi(authResponse.ExpiresIn)
	a.setAuthToken(authResponse.TokenType, authResponse.AccessToken, expiresIn, key, secret)

	return a.token, nil
}
