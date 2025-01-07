/*
Package c2b provides functionality for registering URLs used in the M-Pesa C2B (Customer to Business) API.

This package allows businesses to register validation and confirmation URLs for receiving customer payments. It also includes methods for validating input and decoding responses from M-Pesa API endpoints.

Types and Functions:
- RegisterC2BURLRequest: Represents the request payload for registering a validation and confirmation URL.
- DecodeResponse: Decodes the HTTP response from the M-Pesa API into a structured response or an error.
- FillDefaults: Sets default values for the request parameters.
- Validate: Validates the request parameters for correctness.
*/
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

/*
RegisterC2BURLRequest represents the parameters for registering a C2B validation and confirmation URL.

This request is used to set up a shortcode for accepting customer payments via the M-Pesa API.

Fields:
  - ShortCode (string): The unique M-Pesa shortcode used for business payments.
  - ResponseType (common.ResponseType): Determines how M-Pesa handles unresponsive validation URLs.
    Acceptable values are:
      - `Completed`: Automatically complete the transaction.
      - `Cancelled`: Automatically cancel the transaction.
  - CommandID (common.CommandId): Specifies the command for the request. Defaults to "RegisterURL".
  - ConfirmationURL (string): The URL to receive payment completion notifications.
  - ValidationURL (string): The URL to receive payment validation requests.
*/
type RegisterC2BURLRequest struct {
    ShortCode string `json:"ShortCode"`
    ResponseType common.ResponseType `json:"ResponseType"`
    CommandID common.CommandId `json:"CommandID"`
    ConfirmationURL string `json:"ConfirmationURL"`
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

/*
DecodeResponse decodes the HTTP response from the M-Pesa API.

Parameters:
  - res (*http.Response): The HTTP response from the API.

Returns:
  - (interface{}): A structured `RegisterC2BURLSuccessResponse` or an error.

Behavior:
  - If the response indicates success (ResponseCode "200"), it returns a structured success response.
  - If the response indicates failure, it parses the error details and returns an appropriate error.
*/
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
// decodeError processes errors from the M-Pesa API.
func (s *RegisterC2BURLRequest) decodeError(e registerUrlResponse) error {
    errorCode := e.Header.ResponseCode
    return sdkError.NewSDKError(
        string(errorCode),
        fmt.Sprintf("Url registration failed due to %v", e.Header.ResponseMessage),
        )
}
