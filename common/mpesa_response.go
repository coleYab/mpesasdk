package common;

type MpesaSuccessResponse struct {
    RequestType string
    ConversationID string `json:"ConversationID"`
    OriginatorConversatonId string `json:"OriginatorConversationID"`
    ResponseDescription string `json:"ResponseDescription"`
    ResponseCode string `json:"ResponseCode"`
}

type MpesaErrorResponse struct {
    RequestId string `json:"requestId"`
    ErrorCode string `json:"errorCode"`
    ErrorMessage string `json:"errorMessage"`
}

