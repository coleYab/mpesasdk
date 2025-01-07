package b2c

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
)

// B2CRequest defines the parameters for initiating a Business to Customer (B2C) payment.
// This is used when an organization sends money to a customer.
type B2CRequest struct {
    // InitiatorName is the username of the B2C API operator.
    InitiatorName string `json:"InitiatorName"`

    // SecurityCredential is the encrypted password for the API operator.
    SecurityCredential string `json:"SecurityCredential"`

    // CommandID defines the type of B2C transaction (e.g., SalaryPaymentCommandID).
    CommandID string `json:"CommandID"`

    // Amount is the amount to be sent to the customer.
    Amount uint `json:"Amount"`

    // PartyA is the shortcode from which the money is sent.
    PartyA uint `json:"PartyA"`

    // PartyB is the mobile number of the customer receiving the funds.
    PartyB uint `json:"PartyB"`

    // Remarks are optional comments associated with the transaction.
    Remarks string `json:"Remarks"`

    // QueueTimeOutURL is the URL for notifications if the transaction times out.
    QueueTimeOutURL string `json:"QueueTimeOutURL"`

    // ResultURL is the URL to receive results of the payment request.
    ResultURL string `json:"ResultURL"`

    // Occasion is an optional field for additional transaction details.
    Occasion string `json:"Occasion"`

    OriginatorConversationID string `json:"OriginatorConversationID"`
}

type B2CSuccessResponse  common.MpesaSuccessResponse

func(b *B2CRequest) DecodeResponse(res *http.Response) (interface{}, error) {
    bodyData, _ := io.ReadAll(res.Body)
    responseData := B2CSuccessResponse{}
    err := json.Unmarshal(bodyData, &responseData)
    if err != nil {
        return B2CSuccessResponse{}, err
    }

    if responseData.ResponseCode != "0" {
        errorResponseData := common.MpesaErrorResponse{}
        err := json.Unmarshal(bodyData, &errorResponseData)
        if err != nil {
            return B2CSuccessResponse{}, err
        }
        return B2CSuccessResponse{}, b.decodeError(errorResponseData)
    }

    return responseData, nil
}

func (b *B2CRequest) FillDefaults() {
}

func (b *B2CRequest) Validate() error {
    return nil
}

func (b *B2CRequest) decodeError(e common.MpesaErrorResponse) error {
    return sdkError.NewSDKError(e.ErrorCode, fmt.Sprintf("Request %v failed due to %v", e.RequestId, e.ErrorMessage))
}
