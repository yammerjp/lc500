package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	workerpb "github.com/yammerjp/lc500/worker/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	workerpb.UnimplementedGreetingServiceServer
}

func (s *server) Greet(ctx context.Context, req *workerpb.GreetRequest) (*workerpb.GreetResponse, error) {
	return &workerpb.GreetResponse{Message: "Hello, " + req.GetName()}, nil
}

func (s *server) GreetServerStream(req *workerpb.GreetRequest, stream workerpb.GreetingService_GreetServerStreamServer) error {
	resCount := 5
	for i := 0; i < resCount; i++ {
		if err := stream.Send(&workerpb.GreetResponse{Message: "Hello, " + req.GetName()}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func NewMyServer() *server {
	return &server{}
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer()

	workerpb.RegisterGreetingServiceServer(s, NewMyServer())
	reflection.Register(s)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Printf("failed to serve: %v", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	s.GracefulStop()

	//		iso := v8.NewIsolate()
	//		ctx := v8.NewContext(iso)
	//
	//		scanner := bufio.NewScanner(os.Stdin)
	//		fmt.Print("> ")
	//		for scanner.Scan() {
	//			line := scanner.Text()
	//
	//			script, _ := iso.CompileUnboundScript(line, "main.js", v8.CompileOptions{})
	//			val, err := script.Run(ctx)
	//			if err != nil {
	//				fmt.Printf("ERROR: %v\n", err)
	//				continue
	//			}
	//
	//			fmt.Println(val)
	//
	//			fmt.Print("> ")
	//		}
}
