package utils

import (
	"encoding/base64"
	"fmt"
	"time"
)

// GenerateTimestampAndPassword generates a timestamp and an encrypted password
// required for authenticating M-Pesa API requests.
//
// Parameters:
//   - shortcode: The shortcode used in the M-Pesa API request.
//   - passkey: The passkey provided by M-Pesa for secure transactions.
//
// Returns:
//   - A timestamp in the format "YYYYMMDDHHMMSS".
//   - A base64-encoded password combining the shortcode, passkey, and timestamp.
//
// Example:
//   timestamp, password := GenerateTimestampAndPassword(123456, "myPassKey")
func GenerateTimestampAndPassword(shortcode uint, passkey string) (string, string) {
	timestamp := time.Now().Format("20060102150405")
	password := fmt.Sprintf("%d%s%s", shortcode, passkey, timestamp)
	return timestamp, base64.StdEncoding.EncodeToString([]byte(password))
}
