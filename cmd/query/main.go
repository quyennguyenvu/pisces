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
	"pisces/helper"
	"pisces/storage"
)

func main() {
	config.GetConfig()

	// connect database
	storage.Connect()
	defer storage.Disconnect()

	// connect nats
	handler.Connect()
	defer handler.Disconnect()

	// register logger
	helper.Logging()
	defer helper.CloseFile()

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
	}()

	withServerREST := []rest.OptionFunc{
		rest.WithEventQuery(),
	}

	go func() {
		_ = rest.RunServer(ctx, port.QryGRPCPort, port.QryHTTPPort, withServerREST)
	}()

	withServerGRPC := []grpc.OptionFunc{
		grpc.WithEventQuery(query.NewEventQuery()),
	}
	return grpc.RunServer(ctx, port.QryGRPCPort, withServerGRPC)
}
