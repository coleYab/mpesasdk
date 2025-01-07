package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
)

// AccountBalanceRequest represents the parameters for querying the account balance for a shortcode.
// This is used to check the balance of a business shortcode.
type AccountBalanceRequest struct {
    // CommandID specifies the request type, typically "AccountBalanceCommandID".
    CommandID string `json:"CommandID"`

    // IdentifierType defines the type of identifier used for PartyA (usually "Shortcode").
    IdentifierType int `json:"IdentifierType"` // Changed to int

    // Initiator is the username used to authenticate the balance query request.
    Initiator string `json:"Initiator"`

    // PartyA is the shortcode querying for the balance.
    PartyA int `json:"PartyA"` // Changed to int to match the example

    // QueueTimeOutURL is the URL for notifications if the request times out.
    QueueTimeOutURL string `json:"QueueTimeOutURL"`

    // Remarks are optional comments that can be sent with the request.
    Remarks string `json:"Remarks"`

    // ResultURL is the URL to receive the result of the balance query.
    ResultURL string `json:"ResultURL"`

    // SecurityCredential is the encrypted password for the initiator.
    SecurityCredential string `json:"SecurityCredential"`

    // OriginatorConversationID is a unique identifier for the originator of the transaction.
    OriginatorConversationID string `json:"OriginatorConversationID"`
}

type AccountBalanceSuccessResponse common.MpesaSuccessResponse

func (a *AccountBalanceRequest) DecodeResponse(res *http.Response) (interface{}, error) {
    bodyData, _ := io.ReadAll(res.Body)
    responseData := AccountBalanceSuccessResponse{}
    err := json.Unmarshal(bodyData, &responseData)
    if err != nil {
        return AccountBalanceSuccessResponse{}, err
    }

    if responseData.ResponseCode != "0" {
        errorResponseData := common.MpesaErrorResponse{}
        err := json.Unmarshal(bodyData, &errorResponseData)
        if err != nil {
            return AccountBalanceSuccessResponse{}, err
        }
        return AccountBalanceSuccessResponse{}, a.decodeError(errorResponseData)
    }

    return responseData, nil
}

func (a *AccountBalanceRequest) FillDefaults() {
}

func (a *AccountBalanceRequest) Validate() error {
    return nil
}

func (a *AccountBalanceRequest) decodeError(e common.MpesaErrorResponse) error {
    return sdkError.NewSDKError(e.ErrorCode, fmt.Sprintf("Request %v failed due to %v", e.RequestId, e.ErrorMessage))
}
