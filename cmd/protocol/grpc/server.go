package grpc

import (
	"context"
	"gemini/pb/v1/event"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// OptionFunc ..
type OptionFunc func(*grpc.Server)

// WithEventServer ..
func WithEventServer(eventSrv event.EventServiceServer) OptionFunc {
	return func(srv *grpc.Server) {
		event.RegisterEventServiceServer(srv, eventSrv)
	}
}

// RunServer ..
func RunServer(ctx context.Context, grpcPort string, optFuncs ...OptionFunc) error {
	listen, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	// customer.RegisterCustomerServer(server, cus)
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
