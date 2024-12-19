package main

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

type Reciver struct {
	PhoneNo string
}

type Payment struct {
	AccessToken AccessToken
	Amount      float64
	PaymentDate time.Time
	Reciver     Reciver
}

func getBase64EncodedUserCredential(consumerKey, consumerSecret string) string {
	basicCredential := fmt.Sprintf("%s:%s", consumerKey, consumerSecret)
	return base64.StdEncoding.EncodeToString([]byte(basicCredential))
}

func getAccessToken(consumerKey, consumerSecret string) AccessToken {
	accessToken := AccessToken{}
	url := "https://apisandbox.safaricom.et/v1/token/generate?grant_type=client_credentials"
	method := "GET"
	client := &http.Client{}
	authorizationString := fmt.Sprintf("Bearer %s", getBase64EncodedUserCredential(consumerKey, consumerSecret))
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
	// accessToken.ExpiresAt = time.Now().Add(time.Duration())
	return accessToken
}

func makePayment(accessToken AccessToken) Payment {
	payment := Payment{}

	

	return payment
}

func main() {
	clientKey := "uGkShY9dLosArN02rNSGbcQGFZwVhIgueELktmiNuOJJ0Q0t"
	clientSecret := "lpsECfBqBMMDPAiYOiHuFhiAwPm0G4Jh1cGQLkx9e1kOGOAS0s3CXdwSBJEVIX37"

	accessToken := getAccessToken(clientKey, clientSecret)
	fmt.Println(accessToken)
}
