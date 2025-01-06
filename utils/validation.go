package utils

import (
	"fmt"
	"regexp"
	"strings"
)

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

func ValidateString(toValidate string, minLen, maxLen int) error {
	if minLen > 0 && len(toValidate) < minLen {
        return fmt.Errorf("mpesasdk: validation error: %s is shorter than minLen %v", toValidate, minLen)
	}

	if maxLen > 0 && len(toValidate) > maxLen {
		return fmt.Errorf("mpesasdk: validation error: %s is longer than than maxLen %v", toValidate, maxLen)
	}
	return nil
}
