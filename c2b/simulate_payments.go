package c2b

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
)

// SimulateCustomerInititatedPayment simulates a customer-initiated payment.
// It is used for testing and simulating transactions.
// Only on simulation mode
type SimulateCustomerInititatedPayment struct {
    // Unique Command ID	String	CustomerPayBillOnline
    CommandID string `json:"CommandID"`
    // 	Transaction Amount	String - Numeric	20
    Amount uint64 `json:"Amount"`
    // Phone number of the customer	String - Numeric	0700100100
    Msisdn string `json:"Msisdn"`
    // Unique reference number	String	ET234567
    BillRefNumber string `json:"BillRefNumber"`
    // Unique short code of the organization	String - URL
    ShortCode string `json:"ShortCode"`
}


type SimulatePaymentSuccessResponse  common.MpesaSuccessResponse

func (s *SimulateCustomerInititatedPayment) DecodeResponse(res *http.Response) (SimulatePaymentSuccessResponse, error) {
    bodyData, _ := io.ReadAll(res.Body)
    responseData := SimulatePaymentSuccessResponse{}
    err := json.Unmarshal(bodyData, &responseData)
    if err != nil {
        return SimulatePaymentSuccessResponse{}, err
    }

    if responseData.ResponseCode != "0" {
        errorResponseData := common.MpesaErrorResponse{}
        err := json.Unmarshal(bodyData, &errorResponseData)
        if err != nil {
            return SimulatePaymentSuccessResponse{}, err
        }
        return SimulatePaymentSuccessResponse{}, s.decodeError(errorResponseData)
    }

    return responseData, nil
}

func (s *SimulateCustomerInititatedPayment) FillDefaults() {
}

func (s *SimulateCustomerInititatedPayment) Validate() error {
    return nil
}

func (s *SimulateCustomerInititatedPayment) decodeError(e common.MpesaErrorResponse) error {
    return sdkError.NewSDKError(e.ErrorCode, fmt.Sprintf("Request %v failed with %v", e.RequestId, e.ErrorMessage))
}
