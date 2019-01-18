package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// OptionFunc ..
type OptionFunc func(*grpc.Server)

// RunServer ..
func RunServer(ctx context.Context, grpcPort string, optFuncs ...OptionFunc) error {
	listen, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	for _, optFunc := range optFuncs {
		optFunc(server)
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Printf("starting gRPC server on port %s...\n", grpcPort)
	return server.Serve(listen)
}
