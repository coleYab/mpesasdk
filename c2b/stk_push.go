package c2b

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
	"github.com/coleYab/mpesasdk/utils"
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
    TransactionType common.TransactionType `json:"TransactionType"`

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

    passkey string
}

func (s *STKPushPaymentRequest) SetPasskey(passkey string) {
    s.passkey = passkey
}

type STKPushRequestSuccessResponse struct {
    MerchantRequestID string `json:"MerchantRequestID"`
    CheckoutRequestID string `json:"CheckoutRequestID"`
    ResponseCode string `json:"ResponseCode"`
    ResponseDescription string `json:"ResponseDescription"`
    CustomerMessage string `json:"CustomerMessage"`
}

type STKPushRequestError STKPushRequestSuccessResponse

func (s *STKPushPaymentRequest) DecodeResponse(res *http.Response) (interface{}, error) {
    bodyData, _ := io.ReadAll(res.Body)
    responseData := STKPushRequestSuccessResponse{}
    err := json.Unmarshal(bodyData, &responseData)
    if err != nil {
        return STKPushRequestSuccessResponse{}, err
    }

    switch responseData.ResponseCode {
    case "0":
        return responseData, nil
    case "":
        e := common.MpesaErrorResponse{}
        err := json.Unmarshal(bodyData, &e)
        if err != nil {
            return STKPushRequestSuccessResponse{}, sdkError.ProcessingError(err.Error())
        }
        return STKPushRequestSuccessResponse{}, s.decodeError(e)
    default:
        return STKPushRequestSuccessResponse{}, s.decodeError(common.MpesaErrorResponse{
            RequestId: responseData.MerchantRequestID,
            ErrorCode: responseData.ResponseCode,
            ErrorMessage: responseData.ResponseDescription,
        })
    }

}

func (t *STKPushPaymentRequest) FillDefaults() {
    t.Timestamp, t.Password = utils.GenerateTimestampAndPassword(t.BusinessShortCode, t.passkey)
}

func (t *STKPushPaymentRequest) Validate() error {
    validTransactionTypes := []common.TransactionType{
        common.CustomerBuyGoodsOnlineTransaction,
        common.CustomerPayBillOnlineTransaction,
    }

    if !slices.Contains(validTransactionTypes, t.TransactionType) {
        return sdkError.ValidationError("invalid `TransactionType`")
    }

    if err := utils.ValidateURL(t.CallBackURL); err != nil {
        return err
    }

    return nil
}

func (s *STKPushPaymentRequest) decodeError(e common.MpesaErrorResponse) error {
    errorCode := e.ErrorCode
    switch errorCode {
        case "SVC0403":
            return sdkError.AuthenticationError(e.ErrorMessage)
    }
    return sdkError.CustomError("UNKOWN_ERROR", e.ErrorMessage)
}
