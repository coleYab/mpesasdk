// Package transaction provides functionality for managing M-Pesa transactions,
// including querying transaction status and reversing completed transactions.
package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
)

// TransactionReversalRequest represents the parameters for reversing a transaction.
// It is used to cancel or refund a previously completed transaction.
//
// Fields:
//   - Initiator: The username of the API operator initiating the reversal.
//   - SecurityCredential: The encrypted password for the initiator.
//   - CommandID: Must be "TransactionReversal" for a valid reversal request.
//   - TransactionID: The ID of the transaction to be reversed.
//   - Amount: The amount to be reversed.
//   - ReceiverParty: The organization receiving the reversed funds.
//   - RecieverIdentifierType: Defines the type of organization receiving the transaction.
//   - QueueTimeOutURL: URL for timeout notifications.
//   - ResultURL: URL to receive the result of the reversal.
//   - Remarks: Optional comments about the reversal request.
//   - Occasion: Optional additional details about the transaction.
//   - OriginatorConversationID: Unique identifier for the originating transaction.
type TransactionReversalRequest struct {
	Initiator               string                   `json:"Initiator"`
	SecurityCredential      string                   `json:"SecurityCredential"`
	CommandID               common.CommandId         `json:"CommandID"`
	TransactionID           string                   `json:"TransactionID"`
	Amount                  uint64                   `json:"Amount"`
	ReceiverParty           string                   `json:"ReceiverParty"`
	RecieverIdentifierType  common.IdentifierType    `json:"RecieverIdentifierType"`
	QueueTimeOutURL         string                   `json:"QueueTimeOutURL"`
	ResultURL               string                   `json:"ResultURL"`
	Remarks                 string                   `json:"Remarks"`
	Occasion                string                   `json:"Occasion"`
	OriginatorConversationID string                  `json:"OriginatorConversationID"`
}

// TransactionReversalSuccessResponse represents a successful response for a reversal request.
type TransactionReversalSuccessResponse common.MpesaSuccessResponse

// DecodeResponse decodes the HTTP response for a transaction reversal request.
//
// Parameters:
//   - res: The HTTP response object.
//
// Returns:
//   - An instance of TransactionReversalSuccessResponse if the reversal was successful.
//   - An error if the response indicates a failure or the decoding fails.
func (t *TransactionReversalRequest) DecodeResponse(res *http.Response) (interface{}, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := TransactionReversalSuccessResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	if err != nil {
		return TransactionReversalSuccessResponse{}, err
	}

	if responseData.ResponseCode != "0" {
		errorResponseData := common.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &errorResponseData)
		if err != nil {
			return TransactionReversalSuccessResponse{}, err
		}
		return TransactionReversalSuccessResponse{}, t.decodeError(errorResponseData)
	}

	return responseData, nil
}

// FillDefaults is a placeholder for initializing default values in TransactionReversalRequest.
func (t *TransactionReversalRequest) FillDefaults() {}

// Validate checks the validity of the TransactionReversalRequest parameters.
//
// Returns:
//   - An error if validation fails, or nil if the request is valid.
func (t *TransactionReversalRequest) Validate() error {
	return nil
}

// decodeError processes an M-Pesa error response and returns a structured error.
//
// Parameters:
//   - e: The MpesaErrorResponse containing error details.
//
// Returns:
//   - An error describing the failure with details from the response.
func (t *TransactionReversalRequest) decodeError(e common.MpesaErrorResponse) error {
	return sdkError.NewSDKError(e.ErrorCode, fmt.Sprintf("Request %v failed due to %v", e.RequestId, e.ErrorMessage))
}

