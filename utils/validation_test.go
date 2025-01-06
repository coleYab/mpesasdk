package utils_test

import (
	"log"
	"testing"

	utils "github.com/coleYab/mpesasdk/utils"
)

func TestValidateString(t *testing.T) {
	testCases := map[string][]int{
		"Value 1": {1, 2},
		"Value 2": {-1, 4},
		"Value 3": {1, 8},
		"Value 4": {8, 10},
	}

	expectErr := []bool{
		true, true, false, true,
	}

	idx := 0
	for k, v := range testCases {
		err := utils.ValidateString(k, v[0], v[1])

		if !expectErr[idx] && err != nil {
			log.Fatalf("Test %v: Expecting no errors but got: %v", idx, err.Error())
		}
		if expectErr[idx] && err == nil {
			log.Fatalf("Test %v: Expecting errors but got nil insted", idx)
		}
		idx += 1
	}

}

// utils.ValidateURL(rawURL string)
func TestValidateURL(t *testing.T) {
	validHTTPSUrls := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.github.com",
		"https://www.mozilla.org",
		"https://www.stackoverflow.com",
		"https://192.123.123.12",
	}

	validHTTPUrls := []string{
		"http://www.example.com",
		"http://httpbin.org/get",
		"http://www.wikipedia.org",
		"http://www.example.org",
		"http://www.openai.com",
	}

	invalidUrls := []string{
		"http s://nonexistentwebsite.com",
		"http: //invalid-url",
		"https://l ocalhost:9999",
		"ftp://exam ple.com",
		"htp://www.b adurl.com",
	}

	testCount := 0
	for _, val := range validHTTPSUrls {
		if err := utils.ValidateURL(val); err != nil {
			log.Fatalf("Test %v: Expecting the url to be valid but got: %v", testCount, err.Error())
		}
		testCount += 1
	}

	for _, val := range validHTTPUrls {
		if err := utils.ValidateURL(val); err == nil {
			log.Fatalf("Test %v: Expecting the url to be invalid but got nil insted", testCount)
		}
		testCount += 1
	}
	for _, val := range invalidUrls {
		if err := utils.ValidateURL(val); err == nil {
			log.Fatalf("Test %v: Expecting the url to be invalid but got nil insted", testCount)
		}
		testCount += 1
	}
}

func TestValidateEthiopianPhoneNumber(t *testing.T) {
	phoneNumbers := map[string]bool{
		"251711234567":  true,  // valid number
		"251723456789":  true,  // valid number
		"251700000000":  true,  // valid number
		"251900000000":  false, // invalid number
		"251123456789":  false, // invalid number (incorrect format)
		"2517 123 4567": false, // invalid number (spaces included)
		"252711234567":  false, // invalid number (wrong country code)
		"2511234567":    false, // invalid number (too short)
	}

	for phone, isValid := range phoneNumbers {
		err := utils.ValidateEthiopianPhoneNumber(phone)
		if !isValid && err != nil {
            log.Fatalf("Error: expecing error bu)
		}
		if isValid && err == nil {
		}
	}
}
