// Package b2c provides functionality for initiating and handling Business-to-Customer (B2C) payment requests.
// B2C payments are used when a business sends funds directly to a customer's mobile wallet.
package b2c

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
	"github.com/coleYab/mpesasdk/utils"
)

// B2CRequest defines the parameters required to initiate a B2C payment request.
// This struct is used to send money from a business account to a customer's mobile wallet.
//
// Fields:
//   - InitiatorName: The username of the API operator initiating the request.
//   - SecurityCredential: The encrypted password for the initiator.
//   - CommandID: The type of transaction (e.g., BusinessPayment, SalaryPayment, or PromotionPayment).
//   - Amount: The amount to be sent to the customer.
//   - PartyA: The shortcode of the sender (business).
//   - PartyB: The customer's mobile number (recipient).
//   - Remarks: Optional comments about the transaction.
//   - QueueTimeOutURL: URL to receive notifications if the request times out.
//   - ResultURL: URL to receive the result of the transaction.
//   - Occasion: Optional additional transaction details.
//   - OriginatorConversationID: Unique identifier for the originating transaction.
type B2CRequest struct {
	InitiatorName            string           `json:"InitiatorName"`
	SecurityCredential       string           `json:"SecurityCredential"`
	CommandID                common.CommandId `json:"CommandID"`
	Amount                   uint             `json:"Amount"`
	PartyA                   uint             `json:"PartyA"`
	PartyB                   uint             `json:"PartyB"`
	Remarks                  string           `json:"Remarks"`
	QueueTimeOutURL          string           `json:"QueueTimeOutURL"`
	ResultURL                string           `json:"ResultURL"`
	Occasion                 string           `json:"Occasion"`
	OriginatorConversationID string           `json:"OriginatorConversationID"`
}

// B2CSuccessResponse represents a successful response from the B2C payment API.
type B2CSuccessResponse common.MpesaSuccessResponse

// DecodeResponse processes the HTTP response for a B2C payment request and decodes it into the appropriate response type.
func (b *B2CRequest) DecodeResponse(res *http.Response) (interface{}, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := B2CSuccessResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	if err != nil {
		return B2CSuccessResponse{}, sdkError.ProcessingError(err.Error())
	}

	if responseData.ResponseCode != "0" {
		errorResponseData := common.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &errorResponseData)
		if err != nil {
			return B2CSuccessResponse{}, sdkError.ProcessingError(err.Error())
		}
		return B2CSuccessResponse{}, b.decodeError(errorResponseData)
	}

	return responseData, nil
}

// FillDefaults sets default values for the B2CRequest instance.
// This function is currently a placeholder and can be used to initialize default values in future implementations.
func (b *B2CRequest) FillDefaults() {}

// Validate checks the validity of the B2CRequest parameters.
func (b *B2CRequest) Validate() error {
	validCommands := []common.CommandId{
		common.BusinessPaymentCommand,
		common.SalaryPaymentCommand,
		common.PromotionPaymentCommand,
	}
	if !slices.Contains(validCommands, b.CommandID) {
		return sdkError.ValidationError("unknown CommandID " + string(b.CommandID))
	}

	if err := utils.ValidateURL(b.QueueTimeOutURL); err != nil {
		return err
	}

	if err := utils.ValidateURL(b.ResultURL); err != nil {
		return err
	}

	return nil
}

// decodeError converts a MpesaErrorResponse into a structured error message.
func (b *B2CRequest) decodeError(e common.MpesaErrorResponse) error {
	return sdkError.NewSDKError(e.ErrorCode, fmt.Sprintf("Request %v failed due to %v", e.RequestId, e.ErrorMessage))
}

