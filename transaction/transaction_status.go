package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
)

// TransactionStatusRequest represents the parameters for querying the status of a transaction.
// It is used to retrieve the current status of a previously initiated transaction.
//
// Fields:
//   - CommandID: Must be "TransactionStatusQueryCommandID" for a valid request.
//   - IdentifierType: Type of identifier used for PartyA (e.g., "Shortcode").
//   - Initiator: The username used to authenticate the request.
//   - Occasion: Optional field for additional transaction details.
//   - OriginatorConversationID: Unique identifier for the originating transaction.
//   - PartyA: The shortcode or MSISDN receiving the transaction.
//   - QueueTimeOutURL: URL for timeout notifications.
//   - Remarks: Optional comments about the request.
//   - ResultURL: URL to receive the status query result.
//   - SecurityCredential: Encrypted password for the initiator.
//   - TransactionID: The unique identifier for the M-Pesa transaction.
type TransactionStatusRequest struct {
	CommandID                common.CommandId    `json:"CommandID"`
	IdentifierType           common.IdentifierType `json:"IdentifierType"`
	Initiator                string               `json:"Initiator"`
	Occasion                 string               `json:"Occasion"`
	OriginatorConversationID string               `json:"OriginatorConversationID,omitempty"`
	PartyA                   string               `json:"PartyA"`
	QueueTimeOutURL          string               `json:"QueueTimeOutURL"`
	Remarks                  string               `json:"Remarks"`
	ResultURL                string               `json:"ResultURL"`
	SecurityCredential       string               `json:"SecurityCredential"`
	TransactionID            string               `json:"TransactionID"`
}

// TransactionStatusSuccessResponse represents a successful response for a status query.
type TransactionStatusSuccessResponse common.MpesaSuccessResponse

// DecodeResponse decodes the HTTP response for a transaction status query.
func (t *TransactionStatusRequest) DecodeResponse(res *http.Response) (interface{}, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := TransactionStatusSuccessResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	if err != nil {
		return TransactionStatusSuccessResponse{}, err
	}

	if responseData.ResponseCode != "0" {
		errorResponseData := common.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &errorResponseData)
		if err != nil {
			return TransactionStatusSuccessResponse{}, err
		}
		return TransactionStatusSuccessResponse{}, t.decodeError(errorResponseData)
	}

	return responseData, nil
}

// FillDefaults initializes default values for the TransactionStatusRequest.
func (t *TransactionStatusRequest) FillDefaults() {
	t.CommandID = common.TransactionStatusCommand
}

// Validate checks the validity of the TransactionStatusRequest parameters.
func (t *TransactionStatusRequest) Validate() error {
	return nil
}

// decodeError processes an M-Pesa error response and returns a structured error.
func (t *TransactionStatusRequest) decodeError(e common.MpesaErrorResponse) error {
	return sdkError.NewSDKError(e.ErrorCode, fmt.Sprintf("Request %v failed due to %v", e.RequestId, e.ErrorMessage))
}
