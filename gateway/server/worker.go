package server

import (
	"bytes"
	"context"
	"io"
	"net/http"

	workerapi "github.com/yammerjp/lc500/proto/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerClient struct {
	cc     *grpc.ClientConn
	Client workerapi.WorkerClient
}

func NewWorkerClient(target string, isInsecure bool) (*WorkerClient, error) {
	dialOptions := []grpc.DialOption{}
	if isInsecure {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	cc, err := grpc.NewClient(target, dialOptions...)
	if err != nil {
		return nil, err
	}

	return &WorkerClient{
		cc:     cc,
		Client: workerapi.NewWorkerClient(cc),
	}, nil
}

func (c *WorkerClient) Close() {
	c.cc.Close()
}

type WorkerRequest struct {
	Client *WorkerClient
	Vmid   string
	Req    *http.Request
	Ctx    context.Context
}

func (c *WorkerClient) NewWorkerRequest(ctx context.Context, req *http.Request) (*WorkerRequest, error) {
	res, err := c.Client.InitVM(ctx, &workerapi.InitVMRequest{})
	if err != nil {
		return nil, err
	}

	return &WorkerRequest{
		Client: c,
		Vmid:   res.Vmid,
		Req:    req,
		Ctx:    ctx,
	}, nil
}

func (w *WorkerRequest) Compile(script string) error {
	_, err := w.Client.Client.Compile(w.Ctx, &workerapi.CompileRequest{
		Vmid:   w.Vmid,
		Script: script,
	})
	return err
}

func (w *WorkerRequest) Run(additionalContext string) (*WorkerResponse, error) {
	headers := make(map[string][]string)
	for k, v := range w.Req.Header {
		headers[k] = v
	}
	bodyAll, err := io.ReadAll(w.Req.Body)
	if err != nil {
		return nil, err
	}
	w.Req.Body = io.NopCloser(bytes.NewReader(bodyAll))
	workerHeaders := make(map[string]*workerapi.HeaderValue)
	for k, v := range headers {
		workerHeaders[k] = &workerapi.HeaderValue{Values: v}
	}
	req := workerapi.SetContextRequest{
		Vmid:               w.Vmid,
		HttpRequestMethod:  w.Req.Method,
		HttpRequestHeaders: workerHeaders,
		HttpRequestBody:    string(bodyAll),
		HttpRequestUrl:     w.Req.URL.String(),
		AdditionalContext:  additionalContext,
	}
	if _, err = w.Client.Client.SetContext(w.Ctx, &req); err != nil {
		return nil, err
	}

	resRun, err := w.Client.Client.Run(w.Ctx, &workerapi.RunRequest{
		Vmid:    w.Vmid,
		Dispose: true,
	})
	if err != nil {
		return nil, err
	}
	return &WorkerResponse{RunResponse: resRun}, nil
}

type WorkerResponse struct {
	*workerapi.RunResponse
}

func (w *WorkerResponse) WriteTo(rw http.ResponseWriter) {
	rw.WriteHeader(int(w.HttpResponseStatusCode))
	for k, v := range w.HttpResponseHeaders {
		for _, value := range v.Values {
			rw.Header().Add(k, value)
		}
	}
	rw.Write([]byte(w.HttpResponseBody))
}
