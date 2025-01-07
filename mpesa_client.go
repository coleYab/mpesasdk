// Package mpesasdk provides a comprehensive SDK for interacting with the M-Pesa API,
// allowing developers to integrate mobile money payment services into their applications.
// The package simplifies common M-Pesa operations such as C2B (Customer-to-Business) payments,
// B2C (Business-to-Customer) disbursements, transaction status checks, account balance inquiries,
// and transaction reversals.
//
// The package is designed to abstract the complexities of working directly with M-Pesa APIs by
// providing strongly-typed request/response structures, reusable client configurations, and automatic
// request validation and default value population.
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
)

// MpesaClient is the main client for interacting with the M-Pesa API.
// It encapsulates the necessary configuration, such as authentication credentials,
// environment (sandbox or production), HTTP client, and logging capabilities.
type MpesaClient struct {
    consumerKey    string
    consumerSecret string
    env            common.Enviroment
    client         *client.HttpClient
    logger         *service.Logger
}

// NewMpesaClient creates a new instance of MpesaClient.
//
// Parameters:
//   - consumerKey: Your M-Pesa API consumer key.
//   - consumerSecret: Your M-Pesa API consumer secret.
//   - env: The environment for the M-Pesa API (common.PRODUCTION or common.SANDBOX).
//   - logLevel: Log level for the client (e.g., Debug, Info, Error).
//   - timeout: Timeout for API requests.
//   - maxRetries: Maximum number of retries for failed requests.
//
// Returns:
//   - A pointer to an initialized MpesaClient instance.
//   - An error if any of the input parameters are invalid.
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

    logger.Info("Successfully created mpesa client.")

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
    m.logger.Info("Sending request to %v", endpoint)
    if err := req.Validate(); err != nil {
        m.logger.Error("Request to %v validation failed", endpoint)
        return *new(T), err
    }

    // Populate defaults
    req.FillDefaults()

    response, err := m.client.ApiRequest(m.env, endpoint, method, req, authType)
    if err != nil {
        m.logger.Error("Request to %v api request failed", endpoint)
        return *new(T), err
    }
    defer response.Body.Close()

    // Decode the response and type assert the response failing is impossible
    res, err := req.DecodeResponse(response)
    if err != nil {
        m.logger.Error("Request to %v failed to decode response", endpoint)
    } else {
        m.logger.Info("Request to %v successful", endpoint)
    }
    castedResponse, _ := res.(T)
    return castedResponse, err
}

// RegisterNewURL registers a new URL for receiving C2B (Customer-to-Business) payment notifications.
//
// Parameters:
//   - req: A RegisterC2BURLRequest containing the details of the URL to register.
//
// Returns:
//   - A RegisterC2BURLSuccessResponse if the registration is successful.
//   - An error if the request fails validation or the API call fails.
func (m *MpesaClient) RegisterNewURL(req c2b.RegisterC2BURLRequest) (c2b.RegisterC2BURLSuccessResponse, error) {
    endpoint := "/v1/c2b-register-url/register?apikey=" + m.consumerKey
    return executeRequest[c2b.RegisterC2BURLSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeNone)
}

// MakeB2CPaymentRequest initiates a B2C (Business-to-Customer) payment request.
//
// Parameters:
//   - req: A B2CRequest containing the details of the payment.
//
// Returns:
//   - A B2CSuccessResponse if the payment is successful.
//   - An error if the request fails validation or the API call fails.
func (m *MpesaClient) MakeB2CPaymentRequest(req b2c.B2CRequest) (b2c.B2CSuccessResponse, error) {
    endpoint := "/mpesa/b2c/v2/paymentrequest"
    return executeRequest[b2c.B2CSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}


// SimulateCustomerInitiatedPayment simulates a C2B (Customer-to-Business) payment for testing purposes.
//
// Parameters:
//   - req: A SimulateCustomerInititatedPayment containing the details of the simulated payment.
//
// Returns:
//   - A SimulatePaymentSuccessResponse if the simulation is successful.
//   - An error if the request fails validation or the API call fails.
func (m *MpesaClient) SimulateCustomerInitiatedPayment(req c2b.SimulateCustomerInititatedPayment) (c2b.SimulatePaymentSuccessResponse, error) {
    endpoint := "/mpesa/b2c/simulatetransaction/v1/request"
    return executeRequest[c2b.SimulatePaymentSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// CheckTransactionStatus checks the status of a specific transaction.
//
// Parameters:
//   - req: A TransactionStatusRequest containing the transaction ID and other relevant details.
//
// Returns:
//   - A TransactionStatusSuccessResponse if the transaction status is successfully retrieved.
//   - An error if the request fails validation or the API call fails.
func (m *MpesaClient) CheckTransactionStatus(req transaction.TransactionStatusRequest) (transaction.TransactionStatusSuccessResponse, error) {
    endpoint := "/mpesa/transactionstatus/v1/query"
    return executeRequest[transaction.TransactionStatusSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// AccountBalance retrieves the balance of an account linked to the M-Pesa system.
//
// Parameters:
//   - req: An AccountBalanceRequest containing the details of the account.
//
// Returns:
//   - An AccountBalanceSuccessResponse if the balance is successfully retrieved.
//   - An error if the request fails validation or the API call fails.
func (m *MpesaClient) AccountBalance(req account.AccountBalanceRequest) (account.AccountBalanceSuccessResponse, error) {
    endpoint := "/mpesa/accountbalance/v1/query"
    return executeRequest[account.AccountBalanceSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// STKPushPaymentRequest initiates an STK Push request to facilitate a C2B payment.
//
// Parameters:
//   - passkey: The STK passkey used for the request.
//   - req: An STKPushPaymentRequest containing the payment details.
//
// Returns:
//   - An STKPushRequestSuccessResponse if the payment is successfully initiated.
//   - An error if the request fails validation or the API call fails.
func (m *MpesaClient) STKPushPaymentRequest(passkey string, req c2b.STKPushPaymentRequest) (c2b.STKPushRequestSuccessResponse, error) {
    req.SetPasskey(passkey)
    endpoint := "/mpesa/stkpush/v1/processrequest"
    return executeRequest[c2b.STKPushRequestSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}


// ReverseTransaction reverses a previously completed M-Pesa transaction.
//
// Parameters:
//   - req: A TransactionReversalRequest containing the transaction details to be reversed.
//
// Returns:
//   - A TransactionReversalSuccessResponse if the transaction is successfully reversed.
//   - An error if the request fails validation or the API call fails.
func (m *MpesaClient) ReverseTransaction(req transaction.TransactionReversalRequest) (transaction.TransactionReversalSuccessResponse, error) {
    endpoint := "/mpesa/reversal/v1/request"
    return executeRequest[transaction.TransactionReversalSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

