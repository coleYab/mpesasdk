package validation

import (
	"fmt"
	"net/url"
)

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
