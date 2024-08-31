package server

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

type BlueprintFetcher struct {
	endpoint string
}

func (f *BlueprintFetcher) Host() string {
	blueprintUrl, err := url.Parse(f.endpoint)
	if err != nil {
		return ""
	}
	return blueprintUrl.Host
}

func (f *BlueprintFetcher) Scheme() string {
	blueprintUrl, err := url.Parse(f.endpoint)
	if err != nil {
		return ""
	}
	return blueprintUrl.Scheme
}

func (f *BlueprintFetcher) NewBlueprintRequest(r *http.Request) (*http.Request, error) {
	blueprintUrl, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}
	blueprintUrl.Host = f.Host()
	blueprintUrl.Scheme = f.Scheme()

	bluePrintRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body.Close()

	originalRequestBody := make([]byte, len(bluePrintRequestBody))
	copy(originalRequestBody, bluePrintRequestBody)
	r.Body = io.NopCloser(bytes.NewBuffer(originalRequestBody))

	req, err := http.NewRequest(r.Method, blueprintUrl.String(), io.NopCloser(bytes.NewBuffer(bluePrintRequestBody)))
	if err != nil {
		return nil, err
	}
	req.Header = r.Header.Clone()
	req.Header.Set("Host", r.Host)
	return req, nil
}

func (f *BlueprintFetcher) Fetch(r *http.Request) (*http.Response, error) {
	req, err := f.NewBlueprintRequest(r)
	if err != nil {
		return nil, err
	}
	// redirect no follow
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return client.Do(req)
}

func NewtFetcher(endpoint string) *BlueprintFetcher {
	return &BlueprintFetcher{endpoint: endpoint}
}
