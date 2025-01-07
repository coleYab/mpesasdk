package common

import "net/http"

type MpesaRequest interface {
    DecodeResponse(res *http.Response) (interface{}, error)
    Validate() error
    FillDefaults()
}
