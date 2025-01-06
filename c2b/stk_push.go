package c2b

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    sdkError "github.com/coleYab/mpesasdk/errors"
)

type ReferenceDataRequest struct {
    Key   string `json:"Key"`
    Value string `json:"Value"`
}

// STKPushPaymentRequest: represents the parameters needed to initiate a USSD payment request to M-Pesa.
// Reference: https://developer.safaricom.et/documentation
type STKPushPaymentRequest struct {
    // MerchantRequestID is a unique identifier for the submitted payment request.
    // Example: "SFC-Testing-9146-4216-9455-e3947ac570fc"
    MerchantRequestID string `json:"MerchantRequestID"`

    BusinessShortCode uint `json:"BusinessShortCode"`

    // ReferenceData holds additional details for the transaction.
    // Example: [{"Key": "BundleName", "Value": "Monthly Unlimited Bundle"}]
    ReferenceData []ReferenceDataRequest `json:"ReferenceData"`

    // TransactionType specifies the transaction type (e.g., PayBill or Till numbers).
    TransactionType string `json:"TransactionType"`

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

type STKPushRequestSuccessResponse struct {
    MerchantRequestID string
    CheckoutRequestID string
    ResponseCode string
    ResponseDescription string
    CustomerMessage string
}

type STKPushRequestError STKPushRequestSuccessResponse

func (s *STKPushPaymentRequest) DecodeResponse(res *http.Response) (STKPushRequestSuccessResponse, error) {
    bodyData, _ := io.ReadAll(res.Body)
    responseData := STKPushRequestSuccessResponse{}
    err := json.Unmarshal(bodyData, &responseData)
    if err != nil {
        return STKPushRequestSuccessResponse{}, err
    }

    if responseData.ResponseCode != "0" {
        return STKPushRequestSuccessResponse{}, s.decodeError(STKPushRequestError(responseData))
    }

    return responseData, nil
}

func (a *STKPushPaymentRequest) GetResponseStatus(responseData map[string]interface{}) string {
    return ""
}

func (s *STKPushPaymentRequest) decodeError(e STKPushRequestError) error {
    errorCode := e.ResponseCode
    return sdkError.NewSDKError(
        errorCode,
        fmt.Sprintf("Request %v failed due to %v", e.MerchantRequestID, e.ResponseDescription),
        )
}
