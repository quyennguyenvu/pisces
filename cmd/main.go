package main

import (
	"context"
	"fmt"
	"gemini/cmd/protocol/grpc"
	"gemini/cmd/protocol/rest"
	"gemini/handler"
	"gemini/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	storage.Connect()
	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer ..
func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpPort := os.Getenv("HTTP_PORT")
	grpcPort := os.Getenv("GRPC_PORT")

	ctx := context.Background()

	go func() {
		_ = rest.RunServer(ctx, grpcPort, httpPort)
	}()

	withServer := []grpc.OptionFunc{
		grpc.WithEventServer(handler.NewEventHandler()),
	}

	return grpc.RunServer(ctx, grpcPort, withServer...)
}
