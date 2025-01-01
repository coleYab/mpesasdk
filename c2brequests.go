package mpesasdk;


// RegisterC2BURLRequest represents the parameters for registering a C2B validation and confirmation URL.
// This is used for setting up a shortcode for accepting customer payments.
type RegisterC2BURLRequest struct {
    // ShortCode is the unique M-Pesa shortcode used for business payments.
    ShortCode uint `json:"ShortCode"`

    // ResponseType determines how M-Pesa handles unresponsive validation URLs ("Completed" or "Cancelled").
    ResponseType string `json:"ResponseType"`

    // Use “RegisterURL” to differentiate the service from other services.	String	RegisterURL
    CommandID CommandID `json:"CommandID"`

    // ConfirmationURL is the URL to receive payment completion notifications.
    ConfirmationURL string `json:"ConfirmationURL"`

    // ValidationURL is the URL to receive payment validation requests.
    ValidationURL string `json:"ValidationURL"`
}

type ReferenceDataRequest struct {
    Key string `json:"Key"`
    Value string `json:"Value"`
}

// USSDPushRequest represents the parameters needed to initiate a USSD payment request to M-Pesa.
// Reference: https://developer.safaricom.et/documentation
type USSDPushRequest struct {
    // MerchantRequestID is a unique identifier for the submitted payment request.
    // Example: "SFC-Testing-9146-4216-9455-e3947ac570fc"
    MerchantRequestID string `json:"MerchantRequestID"`

    BusinessShortCode uint `json:"BusinessShortCode"`

    // ReferenceData holds additional details for the transaction.
    // Example: [{"Key": "BundleName", "Value": "Monthly Unlimited Bundle"}]
    ReferenceData []ReferenceDataRequest `json:"ReferenceData"`

    // TransactionType specifies the transaction type (e.g., PayBill or Till numbers).
    TransactionType CommandID `json:"TransactionType"`

    // Password is the base64-encoded password used for encrypting the request.
    Password string `json:"Password"`

    // Timestamp is the time the request is sent, in the format YYYYMMDDHHMMSS.
    Timestamp string `json:"Timestamp"`

    // Amount is the total transaction amount (whole number, no decimals).
    Amount uint64 `json:"Amount"`

    // PartyA is the sender shortcode or business number for the payment.
    PartyA string `json:"PartyA"`

    // PartyB is the recipient shortcode or business number for the payment.
    PartyB string `json:"PartyB"`

    // PhoneNumber is the mobile number of the customer receiving the STK PIN prompt.
    PhoneNumber string `json:"PhoneNumber"`

    // CallBackURL is the URL where M-Pesa will send transaction updates.
    CallBackURL string `json:"CallBackURL"`

    // AccountReference is a unique identifier for the transaction, shown to the customer.
    AccountReference string `json:"AccountReference"`

    // TransactionDesc is a short description for the transaction (1-13 characters).
    TransactionDesc string `json:"TransactionDesc"`
}



// SimulateCustomerInititatedPayment simulates a customer-initiated payment.
// It is used for testing and simulating transactions.
// Only on simulation mode
type SimulateCustomerInititatedPayment struct {
    // Unique Command ID	String	CustomerPayBillOnline
    CommandID CommandID `json:"CommandID"`
    // 	Transaction Amount	String - Numeric	20
    Amount uint64 `json:"Amount"`
    // Phone number of the customer	String - Numeric	0700100100
    Msisdn string `json:"Msisdn"`
    // Unique reference number	String	ET234567
    BillRefNumber string `json:"BillRefNumber"`
    // Unique short code of the organization	String - URL
    ShortCode string `json:"ShortCode"`
}

