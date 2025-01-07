package common;

// MpesaSuccessResponse represents a successful response from the M-Pesa API.
//
// Fields:
//   - RequestType: The type of request that was executed.
//   - ConversationID: A unique identifier for the conversation.
//   - OriginatorConversationID: A unique identifier for the originator of the transaction.
//   - ResponseDescription: A human-readable description of the response.
//   - ResponseCode: The code indicating the success status of the request (typically "0").
type MpesaSuccessResponse struct {
    RequestType             string
    ConversationID          string `json:"ConversationID"`
    OriginatorConversatonId string `json:"OriginatorConversationID"`
    ResponseDescription     string `json:"ResponseDescription"`
    ResponseCode            string `json:"ResponseCode"`
}

// MpesaErrorResponse represents an error response from the M-Pesa API.
//
// Fields:
//   - RequestId: A unique identifier for the request that encountered an error.
//   - ErrorCode: The error code returned by the API.
//   - ErrorMessage: A human-readable description of the error.
type MpesaErrorResponse struct {
    RequestId    string `json:"requestId"`
    ErrorCode    string `json:"errorCode"`
    ErrorMessage string `json:"errorMessage"`
}

