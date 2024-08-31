package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"

	api "github.com/yammerjp/lc500/worker/api"
	"github.com/yammerjp/lc500/worker/pool"
	"github.com/yammerjp/lc500/worker/response"
	"github.com/yammerjp/lc500/worker/vm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	api.UnimplementedWorkerServer
}

func NewServer() *server {
	return &server{}
}

var Pool = pool.NewPool()

func (s *server) InitVM(ctx context.Context, req *api.InitVMRequest) (*api.InitVMResponse, error) {
	vmid := Pool.Generate()
	return &api.InitVMResponse{Vmid: vmid}, nil
}

func (s *server) DisposeVM(ctx context.Context, req *api.DisposeVMRequest) (*api.DisposeVMResponse, error) {
	Pool.Dispose(req.Vmid)
	return &api.DisposeVMResponse{}, nil
}

func (s *server) Compile(ctx context.Context, req *api.CompileRequest) (*api.CompileResponse, error) {
	Pool.CompileScript(req.Vmid, req.Script)
	return &api.CompileResponse{}, nil
}

func (s *server) SetContext(ctx context.Context, req *api.SetContextRequest) (*api.SetContextResponse, error) {
	httpRequest, err := http.NewRequest(req.HttpRequestMethod, req.HttpRequestUrl, strings.NewReader(req.HttpRequestBody))
	if err != nil {
		return nil, err
	}
	vmCtx := vm.NewVMContext(httpRequest, req.AdditionalContext)
	Pool.SetContext(req.Vmid, vmCtx)
	return &api.SetContextResponse{}, nil
}

func (s *server) Run(ctx context.Context, req *api.RunRequest) (*api.RunResponse, error) {
	receiver := response.NewReciever()
	err := Pool.Run(req.Vmid, receiver)
	if err != nil {
		return nil, err
	}
	return receiver.ToGrpc(), nil
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	s := grpc.NewServer()

	api.RegisterWorkerServer(s, NewServer())
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server on port %d", port)
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
	log.Println("gRPC server stopped")
}
