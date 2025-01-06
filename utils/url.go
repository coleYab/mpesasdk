package utils

import (
	"fmt"
	"net/url"
)

func ConstructURL(env, endpoint string) string {
    baseURL := "https://apisandbox.safaricom.et"
    if env == "PRODUCTION" {
        baseURL = "https://api.safaricom.et"
    }
    return fmt.Sprintf("%s%s", baseURL, endpoint)
}

func ValidateURL(rawUrl string) error {
	parsedUrl, err := url.Parse(rawUrl)
    if err != nil {
        return err
    }

    if parsedUrl.Scheme != "https" {
        return fmt.Errorf("Error: url scheme must be https got %v", parsedUrl.Scheme);
    }

	return nil
}
