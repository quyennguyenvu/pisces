package main

import (
	"context"
	"fmt"
	"os"
	"pisces/cmd/protocol/grpc"
	"pisces/cmd/protocol/rest"
	"pisces/config"
	"pisces/handler"
	"pisces/handler/command"
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
		_ = rest.RunServer(ctx, port.CmdGRPCPort, port.CmdHTTPPort)
	}()

	withServerGRPC := []grpc.OptionFunc{
		grpc.WithEventCommand(command.NewEventCommand()),
	}
	return grpc.RunServer(ctx, port.CmdGRPCPort, withServerGRPC...)
}
