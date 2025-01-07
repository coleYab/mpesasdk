package account

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

// AccountBalanceRequest represents the parameters for querying the account balance for a shortcode.
// This is used to check the balance of a business shortcode.
type AccountBalanceRequest struct {
    // CommandID specifies the request type, typically "AccountBalanceCommandID".
    CommandID common.CommandId `json:"CommandID"`

    // IdentifierType defines the type of identifier used for PartyA (usually "Shortcode").
    IdentifierType common.IdentifierType `json:"IdentifierType"`

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
        return AccountBalanceSuccessResponse{}, sdkError.ProcessingError(err.Error())
    }

    if responseData.ResponseCode != "0" {
        errorResponseData := common.MpesaErrorResponse{}
        err := json.Unmarshal(bodyData, &errorResponseData)
        if err != nil {
            return AccountBalanceSuccessResponse{}, sdkError.ProcessingError(err.Error())
        }
        return AccountBalanceSuccessResponse{}, a.decodeError(errorResponseData)
    }

    return responseData, nil
}

func (a *AccountBalanceRequest) FillDefaults() {
    a.CommandID = common.AccountBalanceCommand
}

func (a *AccountBalanceRequest) Validate() error {
    validIdentifiers := []common.IdentifierType{common.MsisdnIdentifierType, common.TillNumberIdentifierType, common.ShortCodeIdentifierType}
    if !slices.Contains(validIdentifiers, a.IdentifierType) {
        return sdkError.ValidationError("unknown identifier type")
    }

    if err := utils.ValidateURL(a.QueueTimeOutURL); err != nil {
        return err
    }

    if err := utils.ValidateURL(a.ResultURL); err != nil {
        return err
    }
    return nil
}

func (a *AccountBalanceRequest) decodeError(e common.MpesaErrorResponse) error {
    return sdkError.NewSDKError("REQUEST_ERROR", fmt.Sprintf("Request %v failed due to %v", e.RequestId, e.ErrorMessage))
}
