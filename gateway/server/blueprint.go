package server

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type BlueprintFetcher struct {
	target string
}

func (f *BlueprintFetcher) NewBlueprintRequest(r *http.Request) (*http.Request, error) {
	blueprintUrl := fmt.Sprintf("http://%s%s", r.Host, r.URL.String())
	slog.Debug("blueprintUrl", "url", blueprintUrl)
	bluePrintRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()

	originalRequestBody := make([]byte, len(bluePrintRequestBody))
	copy(originalRequestBody, bluePrintRequestBody)
	r.Body = io.NopCloser(bytes.NewBuffer(originalRequestBody))

	newReq, err := http.NewRequest(r.Method, blueprintUrl, io.NopCloser(bytes.NewBuffer(bluePrintRequestBody)))
	if err != nil {
		return nil, err
	}

	for k, v := range r.Header {
		for _, value := range v {
			newReq.Header.Add(k, value)
		}
	}
	newReq.Method = r.Method

	return newReq, nil
}

func (f *BlueprintFetcher) Fetch(r *http.Request) (*http.Response, error) {
	req, err := f.NewBlueprintRequest(r)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(fmt.Sprintf("http://%s", f.target))
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
