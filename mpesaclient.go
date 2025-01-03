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
	"github.com/coleYab/mpesasdk/client"
	"github.com/coleYab/mpesasdk/transaction"
	"github.com/coleYab/mpesasdk/validation"
)

type Enviroment string

const (
    EnviromentProduction = "Production"
    EnviromentSandBox = "SandBox"
)

func (e Enviroment) GetBaseUrl() string {
    url := "https://apisandbox.safaricom.et"
    if e == EnviromentProduction {
        url = "https://api.safaricom.et"
    }

    return url
}

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

type App struct {
    initiatorName string
    shortCode uint
	consumerKey    string
	consumerSecret string
    securityCredential string
    passkey string
    authorizationToken AuthorizationToken
    client *client.HttpClient
    enviroment Enviroment
}

func NewMpesaClientRefactored(
    initiatorName string, shortCode uint,
    consumerKey, consumerSecret string, securityCredential string,
    passkey string, timeout uint, maxRetries uint, enviroment Enviroment,
) *App {
    httpClient := client.NewHttpClient(timeout, maxRetries)
    return &App{
        consumerKey: consumerKey,
        consumerSecret: consumerSecret,
        authorizationToken: AuthorizationToken{},
        client: httpClient,
    }
}

type MpesaClient struct {
	consumerKey    string
	consumerSecret string

    authorizationToken AuthorizationToken

    certPath string

    client *http.Client
}


func NewMpesaClient(consumerKey, consumerSecret string) *MpesaClient {
    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    return &MpesaClient{
        consumerKey: consumerKey,
        consumerSecret: consumerSecret,

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

func (m *MpesaClient) makeHttpRequestwithToken(url string, method string, payload []byte, setAuthToken bool) (*http.Response, error) {
    req, err := http.NewRequest(method, url, bytes.NewReader(payload))
    if err != nil {
        return nil, err
    }

    req.Header.Add("Content-Type", "application/json")

    if setAuthToken {
        authToken, err := m.GetAuthorizationToken()
        if err != nil {
            return nil, fmt.Errorf("error getting authorization token: %v", err)
        }

        req.Header.Add("Authorization", authToken)
    }

    res, err := m.client.Do(req)
    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    _, err = io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }

    return res, nil
}

func (m *MpesaClient) RegisterNewURL(req c2b.RegisterC2BURLRequest) (bool, error) {
    url := "https://apisandbox.safaricom.et/v1/c2b-register-url/register?apikey=" + m.consumerKey

    switch req.ResponseType {
    case "Cancelled", "Completed":
        payload, _ := json.Marshal(req);
        res, err := m.makeHttpRequestwithToken(url, "POST", payload, false)
        if err != nil {
            return false, err
        }
        data, _ := io.ReadAll(res.Body)
        fmt.Println(string(data))
        defer res.Body.Close();
        return true, nil
    default:
        return false, fmt.Errorf("invalid ResponseType [%s] provided", req.ResponseType)
    }
}

func (m *MpesaClient) SimulateCustomerInititatedPayment(req c2b.SimulateCustomerInititatedPayment) (bool, error) {
    url :=  "https://apisandbox.safaricom.et/mpesa/b2c/simulatetransaction/v1/request"

    payload, _ := json.Marshal(req)
    res, err := m.makeHttpRequestwithToken(url, "POST", payload, true)
    if err != nil {
        return false, err
    }

    data, _ := io.ReadAll(res.Body)
    fmt.Println(string(data))
    defer res.Body.Close();
    return true, nil
}

func (m *MpesaClient) MakeB2CPaymentRequest(req b2c.B2CRequest) (bool, error) {
     url := "https://apisandbox.safaricom.et/mpesa/b2c/v2/paymentrequest"
     payload, _ := json.Marshal(req)
     res, err := m.makeHttpRequestwithToken(url, "POST", payload, true)
     if err != nil {
         return false, err
     }

    data, _ := io.ReadAll(res.Body)
    fmt.Println(string(data))
    defer res.Body.Close();
    return true, nil
}

func (m *MpesaClient) CheckTransactionStatus(req transaction.TransactionStatusRequest) (bool, error) {
    url :=  "https://apisandbox.safaricom.et/mpesa/transactionstatus/v1/query"

    payload, _ := json.Marshal(req)
    res, err := m.makeHttpRequestwithToken(url, "POST", payload, true)
    if err != nil {
        return false, err
    }

    data, _ := io.ReadAll(res.Body)
    fmt.Println(string(data))
    defer res.Body.Close();
    return true, nil
}

func (m *MpesaClient) AccountBalance(req account.AccountBalanceRequest) (bool, error) {
    url :=  "https://apisandbox.safaricom.et/mpesa/accountbalance/v2/query"

    req.CommandID = "AccountBalance"

    payload, _ := json.Marshal(req)
    res, err := m.makeHttpRequestwithToken(url, "POST", payload, true)
    if err != nil {
        return false, err
    }

    data, _ := io.ReadAll(res.Body)
    fmt.Println(string(data))
    defer res.Body.Close();
    return true, nil
}

func (m *MpesaClient) STKPushPaymentRequest(passkey string, req c2b.USSDPushRequest) (bool, error) {
    url :=  "https://apisandbox.safaricom.et/mpesa/stkpush/v3/processrequest"
    req.Timestamp, req.Password = generateTimestampAndPassword(req.BusinessShortCode, passkey)

    if err := validation.ValidateURL(req.CallBackURL); err != nil {
        return false, err
    }

    payload, _ := json.Marshal(req)
    res, err := m.makeHttpRequestwithToken(url, "POST", payload, true)
    if err != nil {
        return false, err
    }


    data, _ := io.ReadAll(res.Body)
    fmt.Println(string(data))
    defer res.Body.Close();
    return true, nil
}

func generateTimestampAndPassword(shortcode uint, passkey string) (string, string) {
	timestamp := time.Now().Format("20060102150405")
	password := fmt.Sprintf("%d%s%s", shortcode, passkey, timestamp)
	return timestamp, base64.StdEncoding.EncodeToString([]byte(password))
}

func (m *MpesaClient) ReverseTransaction(req transaction.TransactionReversalRequest) (bool, error) {
    url :=  "https://apisandbox.safaricom.et/mpesa/reversal/v2/request"

    if err := validation.ValidateURL(req.ResultURL); err != nil {
        return false, err
    }

    if err := validation.ValidateURL(req.QueueTimeOutURL); err != nil {
        return false, err
    }

    payload, _ := json.Marshal(req)
    res, err := m.makeHttpRequestwithToken(url, "POST", payload, true)
    if err != nil {
        return false, err
    }

    data, _ := io.ReadAll(res.Body)
    fmt.Println(string(data))
    defer res.Body.Close();
    return true, nil
}
