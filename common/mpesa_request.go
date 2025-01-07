// Package common provides shared types and utilities for interacting with the M-Pesa API.
package common

import "net/http"

// MpesaRequest defines the interface that must be implemented by all request types
// to interact with the M-Pesa API.
//
// This interface ensures that all request types provide standard methods for:
//   - Decoding responses from the M-Pesa API.
//   - Validating the request data before sending it to the API.
//   - Populating default values for some fields with known value.
//
// Implementing this interface allows request types to decode success response
// seamlessly with the M-Pesa API client.
//
// Methods:
//
//   DecodeResponse(res *http.Response) (interface{}, error):
//     Decodes the HTTP response from the M-Pesa API into an appropriate Go structure.
//
//   Validate() error:
//     Performs validation checks on the request fields to ensure the data is complete and correct
//     before sending the request to the API. Returns an error if the validation fails.
//
//   FillDefaults():
//     Populates default values for fields in the request. This ensures required fields
//     have valid defaults if not explicitly set by the user.
//
// Example:
//   To create a new request type for an M-Pesa API endpoint, define a struct for the request data
//   and implement the MpesaRequest interface. Here's a simplified example:
//
//   ```go
//   type MyMpesaRequest struct {
//       Field1 string
//       Field2 int
//   }
//
//   func (r *MyMpesaRequest) DecodeResponse(res *http.Response) (interface{}, error) {
//       // Implement response decoding logic
//   }
//
//   func (r *MyMpesaRequest) Validate() error {
//       // Implement validation logic
//   }
//
//   func (r *MyMpesaRequest) FillDefaults() {
//       // Set default values for fields
//   }
//   ```
//
// By adhering to this interface, all request types maintain consistency and can be used interchangeably
// within the M-Pesa SDK.
type MpesaRequest interface {
    // DecodeResponse decodes the HTTP response from the M-Pesa API into an appropriate Go structure.
    //
    // Parameters:
    //   - res: The HTTP response from the M-Pesa API.
    //
    // Returns:
    //   - A decoded response object (as an interface{}) if successful.
    //   - An error if the decoding fails or the response indicates failure.
    DecodeResponse(res *http.Response) (interface{}, error)

    // Validate ensures the request is complete and adheres to the expected format
    // before being sent to the M-Pesa API.
    //
    // Returns:
    //   - nil if the request is valid.
    //   - An error if the request data is incomplete or invalid.
    Validate() error

    // FillDefaults populates default values for optional fields in the request.
    // This ensures that the request has all necessary data before being sent.
    FillDefaults()
}

