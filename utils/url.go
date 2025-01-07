package utils

import (
	"fmt"
	"net/url"

	"github.com/coleYab/mpesasdk/common"
	sdkError "github.com/coleYab/mpesasdk/errors"
)

func ConstructURL(env common.Enviroment, endpoint string) string {
    baseURL := "https://apisandbox.safaricom.et"
    if env == common.PRODUCTION {
        baseURL = "https://api.safaricom.et"
    }
    return fmt.Sprintf("%s%s", baseURL, endpoint)
}

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
