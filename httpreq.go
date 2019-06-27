package httreq

import (
	"net/http"
	"time"
)


var myClient = &http.Client{Timeout: 10 * time.Second}
func Req(url string) (*http.Response, error) {
	res, err := myClient.Get(url)
    if err != nil {
        return nil, err
	}
	return res, nil
}