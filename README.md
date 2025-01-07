# MpesaSdk

`mpesasdk` is a Go SDK for interacting with Safaricom's M-Pesa API. This SDK simplifies integration with M-Pesa's services, enabling operations such as B2C payments, C2B URL registration, STK Push payments, transaction status queries, account balance checks, and transaction reversals.

> **Note**: The SDK is in **active development**. Future updates will include additional features, enhanced error handling, and validation improvements.

## Features

- **B2C Payments**: Transfer funds from a business account to a customer account.
- **C2B URL Registration**: Register URLs for payment notifications.
- **STK Push**: Initiate USSD-based payment requests.
- **Transaction Status**: Query the status of transactions.
- **Account Balance**: Retrieve M-Pesa account balances.
- **Transaction Reversal**: Reverse a completed M-Pesa transaction.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Examples](#examples)
  - [Register C2B URL](#register-c2b-url)
  - [Simulate C2B Payment](#simulate-c2b-payment)
  - [Make B2C Payment](#make-b2c-payment)
  - [Transaction Status Query](#transaction-status-query)
  - [Account Balance Query](#account-balance-query)
  - [Transaction Reversal](#transaction-reversal)
  - [STK Push Payment](#stk-push-payment)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install the SDK, use the following command:

```bash
go get github.com/coleYab/mpesasdk
```

## Quick Start

### Initialize the MpesaClient

```go
client, err := mpesasdk.NewMpesaClient(
    "<consumer_key>",
    "<consumer_secret>",
    common.SANDBOX,
    service.LogLevelInfo,
    10*time.Second,
    3,
)
if err != nil {
    log.Fatalf("Failed to initialize M-Pesa client: %v", err)
}
```

## Examples

### Register C2B URL

```go
response, err := client.RegisterNewURL(c2b.RegisterC2BURLRequest{
    ShortCode:       123456,
    ResponseType:    c2b.ResponseTypeCompleted,
    ConfirmationURL: "https://yourdomain.com/confirmation",
    ValidationURL:   "https://yourdomain.com/validation",
})
```

### Simulate C2B Payment

```go
response, err := client.SimulateCustomerInitiatedPayment(c2b.SimulateCustomerInititatedPayment{
    ShortCode:     "123456",
    BillRefNumber: "INV123",
    Amount:        500,
    Msisdn:        "254700123456",
})
```

### Make B2C Payment

```go
response, err := client.MakeB2CPaymentRequest(b2c.B2CRequest{
    InitiatorName:      "apiuser",
    SecurityCredential: "<security_credential>",
    CommandID:          b2c.BusinessPayment,
    Amount:             1000,
    PartyA:             600000,
    PartyB:             254700123456,
    Remarks:            "Salary Payment",
    QueueTimeOutURL:    "https://yourdomain.com/timeout",
    ResultURL:          "https://yourdomain.com/result",
})
```

### Transaction Status Query

```go
response, err := client.CheckTransactionStatus(transaction.TransactionStatusRequest{
    TransactionID: "LKXXXX1234",
    PartyA:        600000,
    IdentifierType: 1,
    ResultURL:     "https://yourdomain.com/result",
    QueueTimeOutURL: "https://yourdomain.com/timeout",
    Remarks:        "Checking status",
})
```

### Account Balance Query

```go
response, err := client.AccountBalance(account.AccountBalanceRequest{
    PartyA:           600000,
    IdentifierType:   4,
    Remarks:          "Account Balance",
    ResultURL:        "https://yourdomain.com/result",
    QueueTimeOutURL:  "https://yourdomain.com/timeout",
})
```

### Transaction Reversal

```go
response, err := client.ReverseTransaction(transaction.TransactionReversalRequest{
    TransactionID:     "LKXXXX1234",
    PartyA:            600000,
    ReceiverParty:     254700123456,
    IdentifierType:    4,
    Amount:            1000,
    Remarks:           "Reversing transaction",
    ResultURL:         "https://yourdomain.com/result",
    QueueTimeOutURL:   "https://yourdomain.com/timeout",
})
```

### STK Push Payment

```go
response, err := client.STKPushPaymentRequest("<passkey>", c2b.STKPushPaymentRequest{
    BusinessShortCode: 123456,
    Amount:            500,
    PartyA:            "254700123456",
    PartyB:            "123456",
    PhoneNumber:       "254700123456",
    CallBackURL:       "https://yourdomain.com/callback",
    AccountReference:  "INV123",
    TransactionDesc:   "Payment for goods",
})
```

## Contributing

1. Fork the repository.
2. Create a new branch.
3. Commit your changes.
4. Push to the branch.
5. Open a pull request.

## License

This project is licensed under the MIT License.
