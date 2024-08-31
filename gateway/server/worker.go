package server

import (
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

func (w *WorkerRequest) Run(req *http.Request, res *http.Response) (*WorkerResponse, error) {
	workerHeaders := make(map[string]*workerapi.HeaderValue)
	for k, v := range req.Header {
		workerHeaders[k] = &workerapi.HeaderValue{Values: v}
	}
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	setContextReq := &workerapi.SetContextRequest{
		Vmid: w.Vmid,
		HttpRequest: &workerapi.HttpRequest{
			Method:  w.Req.Method,
			Url:     w.Req.URL.String(),
			Headers: workerHeaders,
			Body:    string(requestBody),
		},
		HttpResponse: &workerapi.HttpResponse{
			StatusCode: int32(w.Req.Response.StatusCode),
			Headers:    workerHeaders,
			Body:       string(responseBody),
		},
	}
	if _, err = w.Client.Client.SetContext(w.Ctx, setContextReq); err != nil {
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
	rw.WriteHeader(int(w.HttpResponse.StatusCode))
	for k, v := range w.HttpResponse.Headers {
		for _, value := range v.Values {
			rw.Header().Add(k, value)
		}
	}
	rw.Write([]byte(w.HttpResponse.Body))
}
