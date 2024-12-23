package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AccessToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresAt time.Time
}

// generates a base64 encoing for this specific user credentials that will be used for authorizing
func GetBase64EncodedUserCredential(consumerKey, consumerSecret string) string {
	basicCredential := fmt.Sprintf("%s:%s", consumerKey, consumerSecret)
	return base64.StdEncoding.EncodeToString([]byte(basicCredential))
}

func GetBase64EndodedTransactionPassword(shortCode, clientSecret string, timeStamp time.Time) string {
	// we will need to generate the payment here	

	return shortCode + clientSecret + timeStamp.String()
}

// TODO(coleYab): please try to respond with the valid error messages.
func GetAccessToken(consumerKey, consumerSecret string) AccessToken {
	accessToken := AccessToken{}
	url := "https://apisandbox.safaricom.et/v1/token/generate?grant_type=client_credentials"
	method := "GET"
	client := &http.Client{}
	authorizationString := fmt.Sprintf("Bearer %s", GetBase64EncodedUserCredential(consumerKey, consumerSecret))
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return accessToken
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorizationString)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return accessToken
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return accessToken
	}

	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		fmt.Println(err)
		return accessToken
	}

	// TODO(coleYab): fix the timing by adding expired_at seconds to the time taken
	// accessToken.ExpiresAt = time.Now().Add(time.Duration())
	return accessToken
}
