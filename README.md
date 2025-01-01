
# MpesaSdk 
`
mpesasdk` is a Go SDK for interacting with Safaricom's M-Pesa API. This SDK allows you to easily integrate M-Pesa services into your Go applications. With support for operations like B2C payments, C2B payments, STK push, account balance checks, and transaction status queries, this SDK provides an easy and efficient way to integrate M-Pesa's functionalities.

> **Note**: This SDK is still in **active deployment** and is not yet officially released. There are ongoing improvements planned, including validation, enhanced error handling, and official releases. More tasks will be tackled in upcoming updates.

## Features

- **B2C Payments** - Initiate payments from business accounts to customer accounts.
- **C2B Payments** - Simulate customer-initiated payments to a business account.
- **STK Push** - Request M-Pesa payment via USSD.
- **Transaction Reversal** - Reverse a completed M-Pesa transaction.
- **Transaction Status** - Query the status of a transaction.
- **Account Balance** - Check the balance of an M-Pesa account.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Examples](#examples)
  - [Register New URL](#register-new-url)
  - [Simulate Customer Initiated Payment](#simulate-customer-initiated-payment)
  - [Make B2C Payment](#make-b2c-payment)
  - [Transaction Status Query](#transaction-status-query)
  - [Account Balance](#account-balance)
  - [Transaction Reversal](#transaction-reversal)
  - [STK Push Payment](#stk-push-payment)
- [Contributing](#contributing)
- [License](#license)

## Installation

To get started with the SDK, you need to install it in your Go project. Use the following command:

```bash
go get github.com/coleYab/mpesasdk
```

Make sure you have Go installed and set up on your machine. If not, follow the instructions on [Go's official website](https://golang.org/doc/install).

## Quick Start

To begin using the `mpesasdk`, you need to create an instance of the `MpesaClient` by providing your M-Pesa credentials: `ConsumerKey`, `ConsumerSecret`, `Passkey`, and `SecurityCredential` from [Safaricom Developer Portal](https://developer.safaricom.et/documentation). You can then perform various operations using the client. 

Here’s a basic example:

```go
package main

import (
    "github.com/coleYab/mpesasdk"
)

const consumerKey = "<Your Consumer Key>"
const consumerSecret = "<Your Consumer Secret>"
const passkey = "<Your Passkey>"
const securityCredential = "<Your Security Credentials>"

func main() {
    mpesaClient := mpesasdk.NewMpesaClient(consumerKey, consumerSecret)

    // Register a new C2B URL
    mpesaClient.RegisterNewURL(mpesasdk.RegisterC2BURLRequest{
        ShortCode:       123456,
        ResponseType:    mpesasdk.ResponseTypeCompleted,
        CommandID:       "RegisterURL",
        ConfirmationURL: "https://yourdomain.com/confirmation",
        ValidationURL:   "https://yourdomain.com/validation",
    })
}
```

## Examples

### Register New URL

This example shows how to register a new URL for receiving responses to C2B payments.

```go
mpesaClient.RegisterNewURL(mpesasdk.RegisterC2BURLRequest{
    ShortCode:       123456,
    ResponseType:    mpesasdk.ResponseTypeCompleted,
    CommandID:       "RegisterURL",
    ConfirmationURL: "https://yourdomain.com/confirmation",
    ValidationURL:   "https://yourdomain.com/validation",
})
```

### Simulate Customer Initiated Payment

This example demonstrates how to simulate a customer-initiated payment to a business.

```go
mpesaClient.SimulateCustomerInititatedPayment(mpesasdk.SimulateCustomerInititatedPayment{
    ShortCode:       "123456",
    BillRefNumber:   "ET343434",
    CommandID:       "CustomerPayBillOnline",
    Amount:          54,
    Msisdn:          "251789898989",
})
```

### Make B2C Payment

Initiate a B2C payment to transfer funds from a business account to a customer account.

```go
mpesaClient.MakeB2CPaymentRequest(mpesasdk.B2CRequest{
    InitiatorName:            "testapi",
    SecurityCredential:       securityCredential,
    CommandID:               "BusinessPayment",
    Occasion:                "Payment for services",
    PartyA:                  101010,
    PartyB:                  251700100100,
    Amount:                  100,
    Remarks:                 "Payment for services",
    QueueTimeOutURL:         "https://yourdomain.com/timeout",
    ResultURL:               "https://yourdomain.com/result",
    OriginatorConversationID: "543168755246895",
})
```

### Transaction Status Query

Check the status of a specific transaction.

```go
mpesaClient.CheckTransactionStatus(mpesasdk.TransactionStatusRequest{
    Initiator:              "testapi",
    SecurityCredential:     securityCredential,
    CommandID:              "TransactionStatusQuery",
    TransactionID:          "RHJ4BTOYS8",
    OriginatorConversationID: "AG_20190826_0000777ab7d848b9e721",
    PartyA:                 "101010",
    IdentifierType:         "4",
    ResultURL:              "https://yourdomain.com/api/transaction-status/result",
    QueueTimeOutURL:        "https://yourdomain.com/transaction-status/timeout",
    Remarks:                "OK",
    Occasion:               "OK",
})
```

### Account Balance

Request for account balance.

```go
mpesaClient.AccountBalance(mpesasdk.AccountBalanceRequest{
    CommandID:               mpesasdk.AccountBalanceCommandID,
    Initiator:               "testapi",
    SecurityCredential:      securityCredential,
    PartyA:                  600984,
    IdentifierType:          4,
    Remarks:                 "Balance check",
    ResultURL:               "https://yourdomain.com/result",
    QueueTimeOutURL:         "https://yourdomain.com/timeout",
    OriginatorConversationID: "277406453938994",
})
```

### Transaction Reversal

Reversal of a completed transaction.

```go
mpesaClient.ReverseTransaction(mpesasdk.TransactionReversalRequest{
    Initiator:              "testapiuser",
    SecurityCredential:     "dsam==",
    CommandID:              "TransactionReversal",
    TransactionID:          "A644545RED",
    Amount:                 2000,
    ReceiverParty:          "600610",
    ReceiverIdentifierType: "4",
    ResultURL:              "https://yourdomain.com/result",
    QueueTimeOutURL:        "https://yourdomain.com/timeout",
    Remarks:                "Reversal request",
    Occasion:               "Payment reversal",
    OriginatorConversationID: "543168754246895",
})
```

### STK Push Payment

Initiate a USSD payment request via STK push.

```go
mpesaClient.STKPushPaymentRequest(passkey, mpesasdk.USSDPushRequest{
    MerchantRequestID:  "SFC-Testing-9146-4216-9455-e3947ac570fc",
    BusinessShortCode:  554433,
    Password:           "123",
    Timestamp:          "20160216165627",
    TransactionType:    mpesasdk.CustomerPayBillOnlineCommandID,
    Amount:             10,
    PartyA:             "251700404789",
    PartyB:             "554433",
    PhoneNumber:        "251700404789",
    TransactionDesc:    "Monthly Unlimited Package via Chatbot",
    CallBackURL:        "https://yourdomain.com/callback",
    AccountReference:   "DATA",
    ReferenceData:      []mpesasdk.ReferenceDataRequest{},
})
```

## TODOS

## TODOs

Here is a list of tasks and features planned for the complete development of this SDK:

### 1. **Validation**
- [ ] Implement input validation for API requests.
- [ ] Add validation logic for parameters like `Amount`, `ShortCode`, `TransactionID`, etc., to ensure they meet M-Pesa API standards.

### 2. **Error Handling**
- [ ] Improve error handling across the SDK:
  - [ ] Return more detailed and user-friendly error messages.
  - [ ] Implement retry mechanisms for timeout requests.
  - [ ] Handle different HTTP status codes appropriately (e.g., 400, 500 errors).

### 3. **Official Release**
- [ ] Prepare the SDK for an official release:
  - [ ] Finalize the documentation (e.g., add more examples, clarify usage instructions).
  - [ ] Perform thorough testing to ensure stability and reliability.
  - [ ] Publish the SDK on [GoDoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) for easy reference.
  - [ ] Tag the initial stable version in GitHub (e.g., `v1.0.0`).

### 4. **Unit Tests**
- [ ] Develop a comprehensive suite of unit tests to ensure all API functions work as expected.
- [ ] Include tests for edge cases and invalid inputs.
- [ ] Add tests for error handling scenarios and retries.

### 5. **Integrate OAuth Authentication**
- [x] Implement OAuth 2.0 authentication to securely interact with the M-Pesa API, including handling token expiration and refreshing.

### 6. **Improve Logging**
- [ ] Add detailed logging functionality for tracking SDK usage, requests, responses, and errors.
- [ ] Implement customizable log levels (e.g., `INFO`, `ERROR`, `DEBUG`).

### 7. **Asynchronous Operations**
- [ ] Implement support for asynchronous operations (e.g., background workers) for long-running requests.
- [ ] Provide a mechanism to track the progress of requests that may take time (e.g., B2C payments, balance checks).

### 8. **Error Reporting & Monitoring**
- [ ] Integrate an error reporting system to track unhandled exceptions and errors.

### 9. **Improve Documentation**
- [ ] Add more detailed explanations of various M-Pesa API concepts, such as `ShortCode`, `PartyA`, `PartyB`, and their usage.
- [ ] Add more troubleshooting tips and frequently asked questions (FAQ).
- [ ] Provide examples for integrating with different Go frameworks (e.g., Gin, Echo).

### 10. **Versioning & Compatibility**
- [ ] Implement versioning of the SDK to ensure backward compatibility with future M-Pesa API updates.
- [ ] Add a changelog to document breaking changes and new features with each release.

### 11. **SDK Performance Optimization**
- [ ] Optimize performance for high-volume transactions, ensuring the SDK is scalable and efficient.
- [ ] Minimize network requests where possible and reduce the overall SDK overhead.

### 12. **Release Pipeline & Continuous Integration (CI)**
- [ ] Set up CI/CD pipeline to automatically run tests, build, and release the SDK.
- [ ] Automate versioning and release tagging based on commit history.

## Contributing

We welcome contributions to improve the SDK. If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a pull request.

Please ensure your code adheres to the existing coding style and includes appropriate tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
