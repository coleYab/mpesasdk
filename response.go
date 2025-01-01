package mpesasdk;

type Response struct {
    // MerchantRequestID	This is a global unique Identifier for any submitted payment request.	String	16813-1590513-1
    MerchantRequestID string `json:"MerchantRequestID"`
    // CheckoutRequestID	This is a global unique identifier of the processed checkout transaction request.	String	ws_CO_DMZ_12321_23423476
    CheckoutRequestID string `json:"CheckoutRequestID"`
    // ResponseDescription	Response description is an acknowledgment message from the API that gives the status of the request submission. It usually maps to a specific ResponseCode value. It can be a Success submission message or an error description.	String	-The service request has failed
    // - Invalid Access Token
    // -The service request has been accepted successfully.
    ResponseDescription string `json:"ResponseDescription"`
    // ResponseCode	This is a Numeric status code that indicates the status of the transaction submission. 0 means successful submission and any other code means an error occurred.	Numeric	0 or 404.001.03
    ResponseCode string `json:"ResponseCode"`
    // CustomerMessage	This is a message that your system can display to the customer as an acknowledgment of the payment request submission.	String	E.g.: Success. Request accepted for processing.
    CustomerMessage string `json:"CustomerMessage"`
}
