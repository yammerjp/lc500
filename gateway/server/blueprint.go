package server

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
)

type BlueprintFetcher struct {
	target string
}

func (f *BlueprintFetcher) NewBlueprintRequest(r *http.Request) (*http.Request, error) {
	blueprintUrl := "http://" + r.Host + r.URL.String()
	bluePrintRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()

	originalRequestBody := make([]byte, len(bluePrintRequestBody))
	copy(originalRequestBody, bluePrintRequestBody)
	r.Body = io.NopCloser(bytes.NewBuffer(originalRequestBody))

	return http.NewRequest(r.Method, blueprintUrl, io.NopCloser(bytes.NewBuffer(bluePrintRequestBody)))
}

func (f *BlueprintFetcher) Fetch(r *http.Request) (*http.Response, error) {
	req, err := f.NewBlueprintRequest(r)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("tcp", f.target)
			},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return client.Do(req)
}

func NewBlueprintFetcher(target string) *BlueprintFetcher {
	return &BlueprintFetcher{target: target}
}
