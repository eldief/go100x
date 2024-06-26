package types

import "net/http"

type IHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
