package account

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
	"github.com/stretchr/testify/assert"
)

func TestAccountBalanceRequest_FillDefaults(t *testing.T) {
	req := &AccountBalanceRequest{}
	req.FillDefaults()
	assert.Equal(t, common.AccountBalanceCommand, req.CommandID)
}

func TestAccountBalanceRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request AccountBalanceRequest
		expects error
	}{
		{
			name: "Valid Request",
			request: AccountBalanceRequest{
				IdentifierType:  common.ShortCodeIdentifierType,
				QueueTimeOutURL: "https://example.com/timeout",
				ResultURL:       "https://example.com/result",
			},
			expects: nil,
		},
		{
			name: "Invalid Identifier",
			request: AccountBalanceRequest{
                IdentifierType: common.IdentifierType("999"), // Invalid identifier type
			},
			expects: sdkError.ValidationError("unknown identifier type"),
		},
		{
			name: "Invalid QueueTimeOutURL",
			request: AccountBalanceRequest{
				IdentifierType:  common.ShortCodeIdentifierType,
				QueueTimeOutURL: "invalid-url",
			},
			expects: sdkError.ValidationError("invalid URL"),
		},
		{
			name: "Invalid ResultURL",
			request: AccountBalanceRequest{
				IdentifierType: common.ShortCodeIdentifierType,
				ResultURL:     "invalid-url",
			},
			expects: sdkError.ValidationError("invalid URL"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.request.Validate()
			if test.expects == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expects.Error())
			}
		})
	}
}

func TestAccountBalanceRequest_DecodeResponse(t *testing.T) {
	validResponse := common.MpesaSuccessResponse{
		ResponseCode: "0",
		ResponseDescription: "Success",
	}
	validResponseBody, _ := json.Marshal(validResponse)

	errorResponse := common.MpesaErrorResponse{
		ErrorCode:    "500",
		ErrorMessage: "Internal Server Error",
		RequestId:    "12345",
	}
	errorResponseBody, _ := json.Marshal(errorResponse)

	tests := []struct {
		name        string
		response    *http.Response
		expects     interface{}
		expectError bool
	}{
		{
			name: "Valid Response",
			response: &http.Response{
				Body:       io.NopCloser(io.Reader(bytes.NewReader(validResponseBody))),
				StatusCode: http.StatusOK,
			},
			expects: AccountBalanceSuccessResponse(validResponse),
			expectError: false,
		},
		{
			name: "Error Response",
			response: &http.Response{
				Body:       io.NopCloser(io.Reader(bytes.NewReader(errorResponseBody))),
				StatusCode: http.StatusBadRequest,
			},
			expects: nil,
			expectError: true,
		},
		{
			name: "Malformed Response",
			response: &http.Response{
				Body:       io.NopCloser(strings.NewReader("invalid-json")),
				StatusCode: http.StatusOK,
			},
			expects: nil,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &AccountBalanceRequest{}
			result, err := req.DecodeResponse(test.response)
			if test.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expects, result)
			}
		})
	}
}

func TestAccountBalanceRequest_decodeError(t *testing.T) {
	errorResponse := common.MpesaErrorResponse{
		ErrorCode:    "500",
		ErrorMessage: "Internal Server Error",
		RequestId:    "12345",
	}
	req := &AccountBalanceRequest{}
	err := req.decodeError(errorResponse)
    assert.EqualError(t, err, "REQUEST_ERROR: Request 12345 failed due to Internal Server Error")
}

