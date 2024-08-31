package server

import (
	"bytes"
	"encoding/json"
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

type BlueprintResponse struct {
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

func NewBlueprintResponse(res *http.Response) (*BlueprintResponse, error) {
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	body := string(bodyBytes)
	headers := make(map[string][]string)
	for k, v := range res.Header {
		headers[k] = v
	}

	return &BlueprintResponse{
		Status:  res.StatusCode,
		Headers: headers,
		Body:    body,
	}, nil
}

func (r *BlueprintResponse) ToString() (string, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
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

func (f *BlueprintFetcher) FetchBlueprint(r *http.Request) (string, error) {
	req, err := f.NewBlueprintRequest(r)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	br, err := NewBlueprintResponse(res)
	if err != nil {
		return "", err
	}
	return br.ToString()
}

func NewBlueprintFetcher(endpoint string) *BlueprintFetcher {
	return &BlueprintFetcher{endpoint: endpoint}
}
