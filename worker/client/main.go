package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	workerpb "github.com/yammerjp/lc500/worker/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  workerpb.GreetingServiceClient
)

func main() {
	fmt.Println("start client")
	scanner = bufio.NewScanner(os.Stdin)
	address := "localhost:8080"
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client = workerpb.NewGreetingServiceClient(conn)
	for {
		fmt.Println("1: send request")
		fmt.Println("2: exit")
		fmt.Print("please enter > ")
		scanner.Scan()
		num := scanner.Text()
		switch num {
		case "1":
			Greet(client)
		case "2":
			fmt.Println("bye")
			return
		}
	}
}

func Greet(client workerpb.GreetingServiceClient) {
	fmt.Println("Enter your name:")
	scanner.Scan()
	name := scanner.Text()

	req := &workerpb.GreetRequest{Name: name}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	fmt.Println(res.GetMessage())
}
