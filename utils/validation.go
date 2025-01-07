package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateEthiopianPhoneNumber validates whether a given phone number is a valid Ethiopian Safaricom number.
//
// Criteria for a valid phone number:
//   - Must be 12 digits long.
//   - Must start with "2517" (the country code and prefix for Safaricom numbers).
//   - The remaining 8 digits must be numeric.
//
// Parameters:
//   - phoneNumber: The phone number string to validate.
//
// Returns:
//   - nil if the phone number is valid.
//   - An error if the phone number is invalid.
//
// Example:
//   err := ValidateEthiopianPhoneNumber("251712345678")
//   if err != nil {
//       fmt.Println("Invalid phone number:", err)
//   }
func ValidateEthiopianPhoneNumber(phoneNumber string) error {
	phoneNumber = strings.TrimSpace(phoneNumber)

	if len(phoneNumber) != 12 || !strings.HasPrefix(phoneNumber, "2517") {
        fmt.Printf("Phone no: %v, with len, %v failed here", phoneNumber, len(phoneNumber))
        return fmt.Errorf("mpesasdk: validation error: Invalid Safaricom Phone Number")
	}

	// Ensure the rest of the phone number consists of digits (after '251')
	phonePart := phoneNumber[4:] // Exclude '2517' part
    matches, _ := regexp.MatchString("^[0-9]{8}$", phonePart)
    if !matches {
        fmt.Printf("Phone no: %v, with len, %v failed there", phoneNumber, len(phoneNumber))
        return fmt.Errorf("mpesasdk: validation error: Invalid Safaricom Phone Number")
    }

    return nil
}

// ValidateString checks if a given string meets specified length requirements.
//
// Parameters:
//   - toValidate: The string to validate.
//   - minLen: The minimum allowable length (set to 0 to skip this check).
//   - maxLen: The maximum allowable length (set to 0 to skip this check).
//
// Returns:
//   - nil if the string meets the length requirements.
//   - An error if the string is too short or too long.
//
// Example:
//   err := ValidateString("test", 3, 10)
//   if err != nil {
//       fmt.Println("Validation error:", err)
//   }
func ValidateString(toValidate string, minLen, maxLen int) error {
	if minLen > 0 && len(toValidate) < minLen {
        return fmt.Errorf("mpesasdk: validation error: %s is shorter than minLen %v", toValidate, minLen)
	}

	if maxLen > 0 && len(toValidate) > maxLen {
		return fmt.Errorf("mpesasdk: validation error: %s is longer than than maxLen %v", toValidate, maxLen)
	}
	return nil
}
