package server

import (
	"context"

	api "github.com/yammerjp/lc500/proto/api/v1"
	"github.com/yammerjp/lc500/worker/pool"
	"github.com/yammerjp/lc500/worker/response"
)

type Server struct {
	api.UnimplementedWorkerServer

	vmPool *pool.Pool
}

func NewServer() *Server {
	return &Server{
		vmPool: pool.NewPool(),
	}
}

func (s *Server) InitVM(ctx context.Context, req *api.InitVMRequest) (*api.InitVMResponse, error) {
	vmid := s.vmPool.Generate()
	return &api.InitVMResponse{Vmid: vmid}, nil
}

func (s *Server) DisposeVM(ctx context.Context, req *api.DisposeVMRequest) (*api.DisposeVMResponse, error) {
	s.vmPool.Dispose(req.Vmid)
	return &api.DisposeVMResponse{}, nil
}

func (s *Server) Compile(ctx context.Context, req *api.CompileRequest) (*api.CompileResponse, error) {
	s.vmPool.CompileScript(req.Vmid, req.Script)
	return &api.CompileResponse{}, nil
}

func (s *Server) SetContext(ctx context.Context, req *api.SetContextRequest) (*api.SetContextResponse, error) {
	err := s.vmPool.SetContext(req.Vmid, req)
	if err != nil {
		return nil, err
	}
	return &api.SetContextResponse{}, nil
}

func (s *Server) Run(ctx context.Context, req *api.RunRequest) (*api.RunResponse, error) {
	receiver := response.NewReciever()
	err := s.vmPool.Run(req.Vmid, receiver, req.Dispose)
	if err != nil {
		return nil, err
	}
	return receiver.ToGrpc(), nil
}
