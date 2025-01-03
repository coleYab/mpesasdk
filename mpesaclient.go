package mpesasdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/coleYab/mpesasdk/account"
	"github.com/coleYab/mpesasdk/b2c"
	"github.com/coleYab/mpesasdk/c2b"
	"github.com/coleYab/mpesasdk/transaction"
)

type Enviroment string

const (
    EnviromentProduction = "Production"
    EnviromentSandBox = "SandBox"
)

type AuthorizationToken struct {
    Token string
    CreatedAt time.Time
    ExpiresIn int
}

type AuthResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   string `json:"expires_in"`
    ResultCode  string `json:"resultCode"`
    ResultDesc  string `json:"resultDesc"`
}

const (
    authTypeBearer = "Bearer"
    authTypeBasic = "Basic"
)

type MpesaClient struct {
    consumerKey    string
    consumerSecret string

    authorizationToken AuthorizationToken

    enviroment Enviroment

    client *http.Client
}


func NewMpesaClient(consumerKey, consumerSecret string) *MpesaClient {
    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    return &MpesaClient{
        consumerKey: consumerKey,
        consumerSecret: consumerSecret,
        authorizationToken: AuthorizationToken{},

        client: client,
    }
}

func (m *MpesaClient) setAuthToken(tokenType, token string, expiresIn int) {
    m.authorizationToken.Token = fmt.Sprintf("%v %v", tokenType, token)
    m.authorizationToken.CreatedAt = time.Now()
    m.authorizationToken.ExpiresIn = expiresIn
}

func (m *MpesaClient) GetAuthorizationToken() (string, error) {
    url := "https://apisandbox.safaricom.et/v1/token/generate?grant_type=client_credentials"
    method := "GET"
    if m.authorizationToken.Token != "" && time.Now().Before(m.authorizationToken.CreatedAt.Add(time.Duration(m.authorizationToken.ExpiresIn - 2) * time.Second)) {
        return m.authorizationToken.Token, nil
    }

    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        return "", errors.New("error: while creating auth request")
    }

    req.Header.Add("Content-Type", "application/json")
    req.SetBasicAuth(m.consumerKey, m.consumerSecret)

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

    authResponse := AuthResponse{}
    err = json.Unmarshal(body, &authResponse)
    if err != nil {
        return "", err
    }

    if authResponse.ResultCode != "" {
        return "", fmt.Errorf("Authorization Filed with status code %v due to %v",
            authResponse.ResultCode, authResponse.ResultDesc)
    }

    expiresIn, _ := strconv.Atoi(authResponse.ExpiresIn)
    m.setAuthToken(authResponse.TokenType, authResponse.AccessToken, expiresIn);

    return m.authorizationToken.Token, nil
}

func (m *MpesaClient) RegisterNewURL(req c2b.RegisterC2BURLRequest) (bool, error) {
    endpoint := "/v1/c2b-register-url/register?apikey=" + m.consumerKey

    response, err := m.apiRequest(endpoint, "POST", req, "")
    if err != nil {
       fmt.Printf("ERROR: while making request [%v]\n", err.Error())
        return false, err
    }

    fmt.Println(string(response))
    return true, nil
}

func (m *MpesaClient) SimulateCustomerInititatedPayment(req c2b.SimulateCustomerInititatedPayment) (bool, error) {
    response, err := m.apiRequest("/mpesa/b2c/simulatetransaction/v1/request", "POST", req, authTypeBearer)
    if err != nil {
        return false, err
    }
    fmt.Println(string(response))
    return true, nil
}

func (m *MpesaClient) MakeB2CPaymentRequest(req b2c.B2CRequest) (bool, error) {
    response, err := m.apiRequest("/mpesa/b2c/v2/paymentrequest", "POST", req, authTypeBearer)
    if err != nil {
        return false, err
    }
    fmt.Println(string(response))
    return true, nil
}

func (m *MpesaClient) CheckTransactionStatus(req transaction.TransactionStatusRequest) (bool, error) {
    response, err := m.apiRequest("/mpesa/transactionstatus/v1/query", "POST", req, authTypeBearer)
    if err != nil {
        return false, err
    }
    fmt.Println(string(response))
    return true, nil
}

func (m *MpesaClient) AccountBalance(req account.AccountBalanceRequest) (bool, error) {
    response, err := m.apiRequest("/mpesa/accountbalance/v2/query", "POST", req, authTypeBearer)
    if err != nil {
        return false, err
    }
    fmt.Println(string(response))
    return true, nil
}

func (m *MpesaClient) STKPushPaymentRequest(passkey string, req c2b.USSDPushRequest) (bool, error) {
    req.Timestamp, req.Password = generateTimestampAndPassword(req.BusinessShortCode, passkey)
    response, err := m.apiRequest("/mpesa/stkpush/v3/processrequest", "POST", req, authTypeBearer)
    if err != nil {
        return false, err
    }
    fmt.Println(string(response))
    return true, nil
}


func (m *MpesaClient) constructURL(endpoint string) string {
    baseURL := "https://apisandbox.safaricom.et"
    if m.enviroment == EnviromentProduction {
        baseURL = "https://api.safaricom.et"
    }
    return fmt.Sprintf("%s%s", baseURL, endpoint)
}

func (m *MpesaClient) ReverseTransaction(req transaction.TransactionReversalRequest) (bool, error) {
    response, err := m.apiRequest("/mpesa/reversal/v2/request", "POST", req, authTypeBearer)
    if err != nil {
        return false, err
    }
    fmt.Println(string(response))
    return true, nil
}

func generateTimestampAndPassword(shortcode uint, passkey string) (string, string) {
    timestamp := time.Now().Format("20060102150405")
    password := fmt.Sprintf("%d%s%s", shortcode, passkey, timestamp)
    return timestamp, base64.StdEncoding.EncodeToString([]byte(password))
}

func (m *MpesaClient) apiRequest(endpoint, method string, payload interface{}, authType string) ([]byte, error) {
    url := m.constructURL(endpoint)
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
        authToken, err := m.GetAuthorizationToken()
        if err != nil {
            return nil, err
        }
        req.Header.Add("Authorization", authToken)
    case authTypeBasic:
        req.SetBasicAuth(m.consumerKey, m.consumerSecret)
    }

    res, err := m.client.Do(req)
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

