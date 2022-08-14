package application

import (
	"crypto/tls"
	"net/http"
)

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Status:           "",
		StatusCode:       http.StatusOK,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           map[string][]string{},
		Body:             nil,
		ContentLength:    0,
		TransferEncoding: []string{},
		Close:            false,
		Uncompressed:     false,
		Trailer:          map[string][]string{},
		Request:          req,
		TLS:              &tls.ConnectionState{},
	}, nil
}
