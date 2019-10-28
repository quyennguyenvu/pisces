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

	withServerREST := []rest.OptionFunc{
		rest.WithEventCommand(),
	}

	go func() {
		_ = rest.RunServer(ctx, port.CmdGRPCPort, port.CmdHTTPPort, withServerREST)
	}()

	withServerGRPC := []grpc.OptionFunc{
		grpc.WithEventCommand(command.NewEventCommand()),
	}
	return grpc.RunServer(ctx, port.CmdGRPCPort, withServerGRPC)
}
