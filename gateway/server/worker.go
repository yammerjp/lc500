package server

import (
	"context"
	"io"
	"log/slog"
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
	slog.Info("Running worker request", "url", req.URL.String())
	requestHeaders := make(map[string]*workerapi.HeaderValue)
	for k, v := range req.Header {
		slog.Info("Header", "key", k, "value", v)
		requestHeaders[k] = &workerapi.HeaderValue{Values: v}
	}
	requestBody, err := io.ReadAll(req.Body)
	slog.Info("Request body", "body", string(requestBody))
	if err != nil {
		return nil, err
	}
	responseHeaders := make(map[string]*workerapi.HeaderValue)
	for k, v := range res.Header {
		slog.Info("Header", "key", k, "value", v)
		responseHeaders[k] = &workerapi.HeaderValue{Values: v}
	}
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	slog.Info("Response body", "body", string(responseBody))

	setContextReq := &workerapi.SetContextRequest{
		Vmid: w.Vmid,
		HttpRequest: &workerapi.HttpRequest{
			Method:  req.Method,
			Url:     req.URL.String(),
			Headers: requestHeaders,
			Body:    string(requestBody),
		},
		HttpResponse: &workerapi.HttpResponse{
			StatusCode: int32(res.StatusCode),
			Headers:    responseHeaders,
			Body:       string(responseBody),
		},
	}
	slog.Info("Set context request", "request", setContextReq)

	if _, err = w.Client.Client.SetContext(w.Ctx, setContextReq); err != nil {
		return nil, err
	}
	slog.Info("done Set context request")

	resRun, err := w.Client.Client.Run(w.Ctx, &workerapi.RunRequest{
		Vmid:    w.Vmid,
		Dispose: true,
	})
	if err != nil {
		return nil, err
	}
	slog.Info("Run response", "response", resRun)
	return &WorkerResponse{RunResponse: resRun}, nil
}

type WorkerResponse struct {
	*workerapi.RunResponse
}

func (w *WorkerResponse) WriteTo(rw http.ResponseWriter) {
	for k, v := range w.HttpResponse.Headers {
		for _, value := range v.Values {
			rw.Header().Add(k, value)
		}
	}
	rw.WriteHeader(int(w.HttpResponse.StatusCode))
	rw.Write([]byte(w.HttpResponse.Body))
}
