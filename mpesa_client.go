package mpesasdk

import (
    "errors"
    "net/http"
    "time"

    "github.com/coleYab/mpesasdk/account"
    "github.com/coleYab/mpesasdk/auth"
    "github.com/coleYab/mpesasdk/b2c"
    "github.com/coleYab/mpesasdk/c2b"
    "github.com/coleYab/mpesasdk/client"
    "github.com/coleYab/mpesasdk/common"
    "github.com/coleYab/mpesasdk/service"
    "github.com/coleYab/mpesasdk/transaction"
    "github.com/coleYab/mpesasdk/utils"
)

type MpesaClient struct {
    consumerKey    string
    consumerSecret string
    env            common.Enviroment
    client         *client.HttpClient
    logger         *service.Logger
}

func NewMpesaClient(
    consumerKey, consumerSecret string,
    env common.Enviroment,
    logLevel service.LogLevel,
    timeout time.Duration,
    maxRetries uint,
) (*MpesaClient, error) {
    if env != common.PRODUCTION && env != common.SANDBOX {
        return nil, errors.New("invalid environment: must be either 'Production' or 'Sandbox'")
    }

    if consumerKey == "" || consumerSecret == "" {
        return nil, errors.New("consumer key and consumer secret cannot be empty")
    }

    if timeout <= 0 {
        timeout = 5 * time.Second
    }

    if maxRetries == 0 {
        maxRetries = 1
    }

    auth := auth.NewAuthorizationToken(consumerKey, consumerSecret)
    httpClient := client.NewHttpClient(timeout, maxRetries, auth)
    logger := service.NewLogger(logLevel)

    return &MpesaClient{
        consumerKey:    consumerKey,
        consumerSecret: consumerSecret,
        env:            env,
        client:         httpClient,
        logger:         logger,
    }, nil
}

func executeRequest[T any](m *MpesaClient, req common.MpesaRequest, endpoint, method string, authType string) (T, error) {
    // Validate the request
    if err := req.Validate(); err != nil {
        return *new(T), err
    }

    // Populate defaults
    req.FillDefaults()

    response, err := m.client.ApiRequest(string(m.env), endpoint, method, req, authType)
    if err != nil {
        return *new(T), err
    }
    defer response.Body.Close()

    // Decode the response and type assert the response failing is impossible
    res, err := req.DecodeResponse(response)
    castedResponse, _ := res.(T)
    return castedResponse, err
}

func (m *MpesaClient) RegisterNewURL(req c2b.RegisterC2BURLRequest) (c2b.RegisterC2BURLSuccessResponse, error) {
    endpoint := "/v1/c2b-register-url/register?apikey=" + m.consumerKey
    return executeRequest[c2b.RegisterC2BURLSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeNone)
}

func (m *MpesaClient) MakeB2CPaymentRequest(req b2c.B2CRequest) (b2c.B2CSuccessResponse, error) {
    endpoint := "/mpesa/b2c/v2/paymentrequest"
    return executeRequest[b2c.B2CSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *MpesaClient) SimulateCustomerInitiatedPayment(req c2b.SimulateCustomerInititatedPayment) (c2b.SimulatePaymentSuccessResponse, error) {
    endpoint := "/mpesa/c2b/simulate"
    return executeRequest[c2b.SimulatePaymentSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *MpesaClient) CheckTransactionStatus(req transaction.TransactionStatusRequest) (transaction.TransactionStatusSuccessResponse, error) {
    endpoint := "/mpesa/transactionstatus/v1/query"
    return executeRequest[transaction.TransactionStatusSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *MpesaClient) AccountBalance(req account.AccountBalanceRequest) (account.AccountBalanceSuccessResponse, error) {
    endpoint := "/mpesa/accountbalance/v1/query"
    return executeRequest[account.AccountBalanceSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *MpesaClient) STKPushPaymentRequest(passkey string, req c2b.STKPushPaymentRequest) (c2b.STKPushRequestSuccessResponse, error) {
    // TODO(coleYab): refactor this to be in fill defaults
    req.Timestamp, req.Password = utils.GenerateTimestampAndPassword(req.BusinessShortCode, passkey)
    endpoint := "/mpesa/stkpush/v1/processrequest"
    return executeRequest[c2b.STKPushRequestSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *MpesaClient) ReverseTransaction(req transaction.TransactionReversalRequest) (transaction.TransactionReversalSuccessResponse, error) {
    endpoint := "/mpesa/reversal/v1/request"
    return executeRequest[transaction.TransactionReversalSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

