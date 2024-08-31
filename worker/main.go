package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"

	api "github.com/yammerjp/lc500/proto/api/v1"
	"github.com/yammerjp/lc500/worker/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	s := grpc.NewServer()

	api.RegisterWorkerServer(s, server.NewServer())
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
