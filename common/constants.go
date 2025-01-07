package common

// CommandId represents the type of command being executed in an M-Pesa API request.
//
// Predefined Command IDs:
//   - CustomerPayBillOnlineCommand: Used for PayBill transactions by customers.
//   - AccountBalanceCommand: Used for querying the account balance of a shortcode.
//   - CustomerBuyGoodsOnlineCommand: Used for Buy Goods transactions by customers.
//   - BusinessPaymentCommand: Used for making business-to-business payments.
//   - SalaryPaymentCommand: Used for salary disbursements.
//   - PromotionPaymentCommand: Used for promotional payments.
//   - RegisterURLCommand: Used for registering callback URLs for C2B (Customer-to-Business) transactions.
//   - TransactionStatusCommand: Used for querying the status of a transaction.
//   - TransactionReversalCommand: Used for reversing a transaction.
type CommandId string

const (
    CustomerPayBillOnlineCommand  CommandId = "CustomerPayBillOnline"
    AccountBalanceCommand         CommandId = "AccountBalance"
    CustomerBuyGoodsOnlineCommand CommandId = "CustomerBuyGoodsOnline"
    BusinessPaymentCommand        CommandId = "BusinessPayment"
    SalaryPaymentCommand          CommandId = "SalaryPayment"
    PromotionPaymentCommand       CommandId = "PromotionPayment"
    RegisterURLCommand            CommandId = "RegisterURL"
    TransactionStatusCommand      CommandId = "TransactionStatusQuery"
    TransactionReversalCommand    CommandId = "TransactionReversal"
)

// IdentifierType represents the type of identifier used in M-Pesa API requests.
//
// Predefined Identifier Types:
//   - MsisdnIdentifierType: Represents a mobile number (MSISDN).
//   - TillNumberIdentifierType: Represents a Till number.
//   - ShortCodeIdentifierType: Represents a shortcode.
type IdentifierType string

const (
    MsisdnIdentifierType   IdentifierType = "1"
    TillNumberIdentifierType IdentifierType = "2"
    ShortCodeIdentifierType IdentifierType = "4"
)

// Environment defines the operating environment for the M-Pesa API.
//
// Predefined Environments:
//   - PRODUCTION: Indicates the live production environment.
//   - SANDBOX: Indicates the sandbox environment for testing.
type Enviroment string

const (
    PRODUCTION Enviroment = "Production"
    SANDBOX    Enviroment = "SandBox"
)

// TransactionType represents the type of transaction being executed.
//
// Predefined Transaction Types:
//   - CustomerPayBillOnlineTransaction: Represents PayBill transactions by customers.
//   - CustomerBuyGoodsOnlineTransaction: Represents Buy Goods transactions by customers.
type TransactionType string

const (
    CustomerPayBillOnlineTransaction  TransactionType = "CustomerPayBillOnline"
    CustomerBuyGoodsOnlineTransaction TransactionType = "CustomerBuyGoodsOnline"
)

// ResponseType represents the type of response from the M-Pesa API.
//
// Predefined Response Types:
//   - CompletedResponse: Indicates that the request was successfully completed.
//   - CancelledResponse: Indicates that the request was canceled.
type ResponseType string

const (
    CompletedResponse ResponseType = "Completed"
    CancelledResponse ResponseType = "Cancelled"
)

