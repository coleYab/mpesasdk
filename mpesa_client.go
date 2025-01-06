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

// TODO(coleYab): add suppport for more data validation and pagination
type MpesaClient struct {
    consumerKey         string
    consumerSecret      string
    env                 common.Enviroment
    client              *client.HttpClient
    logger              *service.Logger
}

// This is a function that will create the MpesaClient that will enable the users to
// interact with the mpesa api
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
        timeout = 5
    }

    if maxRetries < 0 {
        maxRetries = 1
    }

    auth := auth.NewAuthorizationToken(consumerKey, consumerSecret);
    httpClient := client.NewHttpClient(timeout, maxRetries, auth)
    logger := service.NewLogger(logLevel)
    return &MpesaClient{
        consumerKey:        consumerKey,
        consumerSecret:     consumerSecret,
        env:                env,
        client:             httpClient,
        logger:             logger,
    }, nil
}

func (m *MpesaClient) RegisterNewURL(req c2b.RegisterC2BURLRequest) (c2b.RegisterC2BURLSuccessResponse, error) {
    endpoint := "/v1/c2b-register-url/register?apikey=" + m.consumerKey
    response, err := m.apiRequest(endpoint, "POST", req, auth.AuthTypeNone)
    if err != nil {
        return c2b.RegisterC2BURLSuccessResponse{}, err
    }

    return req.DecodeResponse(response)
}

func (m *MpesaClient) SimulateCustomerInititatedPayment(req c2b.SimulateCustomerInititatedPayment) (c2b.SimulatePaymentSuccessResponse, error) {
    response, err := m.apiRequest("/mpesa/b2c/simulatetransaction/v1/request", "POST", req, auth.AuthTypeBearer)
    if err != nil {
        return c2b.SimulatePaymentSuccessResponse{}, err
    }

    return req.DecodeResponse(response)
}

func (m *MpesaClient) MakeB2CPaymentRequest(req b2c.B2CRequest) (b2c.B2CSuccessResponse, error) {
    response, err := m.apiRequest("/mpesa/b2c/v2/paymentrequest", "POST", req, auth.AuthTypeBearer)
    if err != nil {
        return b2c.B2CSuccessResponse{}, err
    }

    return req.DecodeResponse(response)
}

func (m *MpesaClient) CheckTransactionStatus(req transaction.TransactionStatusRequest) (transaction.TransactionStatusSuccessResponse, error) {
    response, err := m.apiRequest("/mpesa/transactionstatus/v1/query", "POST", req, auth.AuthTypeBearer)
    if err != nil {
        return transaction.TransactionStatusSuccessResponse{}, err
    }
    return req.DecodeResponse(response)
}

func (m *MpesaClient) AccountBalance(req account.AccountBalanceRequest) (account.AccountBalanceSuccessResponse, error) {
    response, err := m.apiRequest("/mpesa/accountbalance/v2/query", "POST", req, auth.AuthTypeBearer)
    if err != nil {
        return account.AccountBalanceSuccessResponse{}, err
    }
    return req.DecodeResponse(response)
}

func (m *MpesaClient) STKPushPaymentRequest(passkey string, req c2b.STKPushPaymentRequest) (c2b.STKPushRequestSuccessResponse, error) {
    req.Timestamp, req.Password = utils.GenerateTimestampAndPassword(req.BusinessShortCode, passkey)
    response, err := m.apiRequest("/mpesa/stkpush/v3/processrequest", "POST", req, auth.AuthTypeBearer)
    if err != nil {
        return c2b.STKPushRequestSuccessResponse{}, err
    }

    return req.DecodeResponse(response)
}

func (m *MpesaClient) ReverseTransaction(req transaction.TransactionReversalRequest) (transaction.TransactionReversalSuccessResponse, error) {
    response, err := m.apiRequest("/mpesa/reversal/v2/request", "POST", req, auth.AuthTypeBearer)
    if err != nil {
        return transaction.TransactionReversalSuccessResponse{}, err
    }
    return req.DecodeResponse(response)
}

func (m *MpesaClient) apiRequest(endpoint, method string, payload interface{}, authType string) (*http.Response, error) {
    return m.client.ApiRequest(string(m.env), endpoint, method, payload, authType)
}

