package common

import "net/http"

type MpesaRequest[T any] interface {
	DecodeResponse(res *http.Response) (T, error)
    Validate() error
    FillDefaults()
}
