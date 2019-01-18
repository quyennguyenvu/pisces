package main

import (
	"context"
	"fmt"
	"os"
	"pisces/cmd/protocol/grpc"
	"pisces/cmd/protocol/rest"
	"pisces/config"
	"pisces/handler"
	"pisces/handler/query"
	"pisces/handler/subscription"
	"pisces/storage"
)

func main() {
	// connect database
	storage.Connect()

	// connect nats
	handler.Connect()
	defer handler.Disconnect()

	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer ..
func RunServer() error {
	port := config.GetPort()
	ctx := context.Background()

	go func() {
		subscription.Run()
		_ = rest.RunServer(ctx, port.QryGRPCPort, port.QryHTTPPort)
	}()

	withServerGRPC := []grpc.OptionFunc{
		grpc.WithEventQuery(query.NewEventQuery()),
	}
	return grpc.RunServer(ctx, port.QryGRPCPort, withServerGRPC...)
}
