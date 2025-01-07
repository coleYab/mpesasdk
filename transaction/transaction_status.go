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
// It is used to check the status of a previously initiated transaction.
type TransactionStatusRequest struct {
	// CommandID is the type of request (usually "TransactionStatusQueryCommandID").
	CommandID common.CommandId `json:"CommandID"`

	// IdentifierType defines the type of identifier used for PartyA (e.g., "Shortcode").
	IdentifierType common.IdentifierType `json:"IdentifierType"`

	// Initiator is the username used to authenticate the transaction query.
	Initiator string `json:"Initiator"`

	// Occasion is an optional field for additional transaction details.
	Occasion string `json:"Occasion"`

	// OriginatorConversationID is the unique identifier for the transaction request.
	OriginatorConversationID string `json:"OriginatorConversationID,omitempty"`

	// PartyA is the shortcode or MSISDN receiving the transaction.
	PartyA string `json:"PartyA"`

	// QueueTimeOutURL is the URL for notifications if the request times out.
	QueueTimeOutURL string `json:"QueueTimeOutURL"`

	// Remarks are optional comments that will be sent with the transaction.
	Remarks string `json:"Remarks"`

	// ResultURL is the URL to receive the result of the status query.
	ResultURL string `json:"ResultURL"`

	// SecurityCredential is the encrypted password for the initiator to authenticate the request.
	SecurityCredential string `json:"SecurityCredential"`

	// TransactionID is the unique identifier for the M-Pesa transaction.
	TransactionID string `json:"TransactionID"`
}

type TransactionStatusSuccessResponse common.MpesaSuccessResponse

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

func (t *TransactionStatusRequest) FillDefaults() {
    t.CommandID = common.TransactionStatusCommand
}


func (t *TransactionStatusRequest) Validate() error {
	return nil
}

func (t *TransactionStatusRequest) decodeError(e common.MpesaErrorResponse) error {
	return sdkError.NewSDKError(e.ErrorCode, fmt.Sprintf("Request %v failed due to %v", e.RequestId, e.ErrorMessage))
}
