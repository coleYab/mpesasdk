// Package utils provides utility functions for common tasks such as URL validation,
// constructing API URLs, password generation, and validation of input data.
// These utilities are designed to assist with the preparation and validation of requests
// for the M-Pesa SDK.
package utils

import (
	"fmt"
	"net/url"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
)

// ConstructURL constructs the full API URL for the M-Pesa endpoint based on the environment.
//
// Parameters:
//   - env: The environment (e.g., PRODUCTION or SANDBOX).
//   - endpoint: The endpoint path for the desired API resource.
//
// Returns:
//   - A full API URL string.
//
// Example:
//   url := ConstructURL(common.SANDBOX, "/mpesa/accountbalance/v1/query")
func ConstructURL(env common.Enviroment, endpoint string) string {
    baseURL := "https://apisandbox.safaricom.et"
    if env == common.PRODUCTION {
        baseURL = "https://api.safaricom.et"
    }
    return fmt.Sprintf("%s%s", baseURL, endpoint)
}

// ValidateURL validates that the given URL is a properly formatted HTTPS URL.
//
// Parameters:
//   - rawUrl: The URL string to validate.
//
// Returns:
//   - nil if the URL is valid.
//   - An error if the URL is invalid (e.g., parsing failure, non-HTTPS scheme).
//
// Example:
//   err := ValidateURL("https://example.com")
//   if err != nil {
//       fmt.Println("Invalid URL:", err)
//   }
func ValidateURL(rawUrl string) error {
	parsedUrl, err := url.Parse(rawUrl)
    if err != nil {
        return sdkError.ValidationError("url parsing failed" + err.Error())
    }

    if parsedUrl.Scheme != "https" {
        return sdkError.ValidationError("invalid URL scheme expected https")
    }

	return nil
}
