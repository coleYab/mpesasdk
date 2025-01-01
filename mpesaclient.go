package mpesasdk

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"embed"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/coleYab/mpesasdk/validation"
)

type AuthorizationToken struct {
    Token string
    CreatedAt time.Time
    ExpiresIn int
}

const (
    ResponseTypeCompleted string = "Completed"
    ResponseTypeCancelled string = "Cancelled"
)


type AuthResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   string `json:"expires_in"`
    ResultCode  string `json:"resultCode"`
    ResultDesc  string `json:"resultDesc"`
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
        if authToken == "" {
            return nil, errors.New("authorization token is empty")
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

func (m *MpesaClient) RegisterNewURL(req RegisterC2BURLRequest) (bool, error) {
    url := "https://apisandbox.safaricom.et/v1/c2b-register-url/register?apikey=" + m.consumerKey

    switch req.ResponseType {
    case ResponseTypeCancelled, ResponseTypeCompleted:
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

func (m *MpesaClient) SimulateCustomerInititatedPayment(req SimulateCustomerInititatedPayment) (bool, error) {
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

func (m *MpesaClient) MakeB2CPaymentRequest(req B2CRequest) (bool, error) {
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

func (m *MpesaClient) CheckTransactionStatus(req TransactionStatusRequest) (bool, error) {
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

func (m *MpesaClient) AccountBalance(req AccountBalanceRequest) (bool, error) {
    url :=  "https://apisandbox.safaricom.et/mpesa/accountbalance/v2/query"

    req.CommandID = AccountBalanceCommandID

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

func (m *MpesaClient) STKPushPaymentRequest(passkey string, req USSDPushRequest) (bool, error) {
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

//go:embed certs/production.cert
var certFS embed.FS
func (m *MpesaClient) encodeInitiatorPassword(initiatorPassword string) (string, error) {
    m.certPath = "certs/production.cert"
    publicKey, err := certFS.ReadFile(m.certPath)
    if err != nil {
        return "", fmt.Errorf("mpesa: read cert: %v", err)
    }

    block, _ := pem.Decode(publicKey)

    var cert *x509.Certificate
    cert, err = x509.ParseCertificate(block.Bytes)
    if err != nil {
        return "", fmt.Errorf("mpesa:parse cert: %v", err)
    }

    rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
    reader := rand.Reader
    signature, err := rsa.EncryptPKCS1v15(reader, rsaPublicKey, []byte(initiatorPassword))
    if err != nil {
        return "", fmt.Errorf("mpesa: encrypt password: %v", err)
    }

    return base64.StdEncoding.EncodeToString(signature), nil
}

func generateTimestampAndPassword(shortcode uint, passkey string) (string, string) {
	timestamp := time.Now().Format("20060102150405")
	password := fmt.Sprintf("%d%s%s", shortcode, passkey, timestamp)
	return timestamp, base64.StdEncoding.EncodeToString([]byte(password))
}

func (m *MpesaClient) ReverseTransaction(req TransactionReversalRequest) (bool, error) {
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