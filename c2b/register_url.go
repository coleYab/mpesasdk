package c2b

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

// RegisterC2BURLRequest represents the parameters for registering a C2B validation and confirmation URL.
// This is used for setting up a shortcode for accepting customer payments.
type RegisterC2BURLRequest struct {
    // ShortCode is the unique M-Pesa shortcode used for business payments.
    ShortCode string `json:"ShortCode"`

    // ResponseType determines how M-Pesa handles unresponsive validation URLs ("Completed" or "Cancelled").
    ResponseType common.ResponseType `json:"ResponseType"`

    // Use “RegisterURL” to differentiate the service from other services.	String	RegisterURL
    CommandID common.CommandId `json:"CommandID"`

    // ConfirmationURL is the URL to receive payment completion notifications.
    ConfirmationURL string `json:"ConfirmationURL"`

    // ValidationURL is the URL to receive payment validation requests.
    ValidationURL string `json:"ValidationURL"`
}

type registerUrlResponse struct {
    Header struct {
        ResponseCode string `json:"responseCode"`
        ResponseMessage string `json:"responseMessage"`
        CustomerMessage string `json:"customerMessage"`
    } `json:"header"`
}

type RegisterC2BURLSuccessResponse common.MpesaSuccessResponse

func (s *RegisterC2BURLRequest) DecodeResponse(res *http.Response) (interface{}, error) {
    bodyData, _ :=  io.ReadAll(res.Body)
    responseData := registerUrlResponse{}
    err := json.Unmarshal(bodyData, &responseData)
    if err != nil {
        return RegisterC2BURLSuccessResponse{}, sdkError.ProcessingError(err.Error())
    }

    switch responseData.Header.ResponseCode {
    case "200":
        return RegisterC2BURLSuccessResponse{
            ResponseCode: string(responseData.Header.ResponseCode),
            ResponseDescription: responseData.Header.ResponseMessage,
        }, nil
    case "":
        // In this case an error with the http.Response.StatusCode != 200 so we need the defult error handling mechanism
        errorResponse := common.MpesaErrorResponse{}
        err := json.Unmarshal(bodyData, &errorResponse)
        if err != nil {
            return RegisterC2BURLSuccessResponse{}, sdkError.ProcessingError(err.Error())
        }
        responseData.Header.ResponseCode = errorResponse.ErrorCode
        responseData.Header.ResponseMessage = errorResponse.ErrorMessage
        fallthrough // use the default error handling after that to improve consistency
    default:
        return RegisterC2BURLSuccessResponse{}, s.decodeError(responseData)
    }
}

func (t *RegisterC2BURLRequest) FillDefaults() {
    t.CommandID = common.RegisterURLCommand
}

func (t *RegisterC2BURLRequest) Validate() error {
    validResponseType := []common.ResponseType{common.CompletedResponse, common.CancelledResponse}
    if !slices.Contains(validResponseType, t.ResponseType) {
        return sdkError.ValidationError("invalid response type " + string(t.ResponseType))
    }

    if err := utils.ValidateURL(t.ConfirmationURL); err != nil {
        return err
    }

    if err := utils.ValidateURL(t.ValidationURL); err != nil {
        return err
    }

    return nil
}

func (s *RegisterC2BURLRequest) decodeError(e registerUrlResponse) error {
    errorCode := e.Header.ResponseCode
    return sdkError.NewSDKError(
        string(errorCode),
        fmt.Sprintf("Url registration failed due to %v", e.Header.ResponseMessage),
        )
}
