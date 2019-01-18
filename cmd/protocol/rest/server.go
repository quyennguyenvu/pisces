package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// OptionFunc ..
type OptionFunc func(context.Context, string)

var mux *runtime.ServeMux
var dialOpts []grpc.DialOption

// RunServer ..
func RunServer(ctx context.Context, grpcPort, httpPort string, opts ...runtime.ServeMuxOption) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux = runtime.NewServeMux(opts...)
	dialOpts = []grpc.DialOption{grpc.WithInsecure()}

	addServices(ctx, grpcPort)

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Printf("%v\n", c)
		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
	}()

	log.Printf("starting HTTP/REST gateway on port %s...\n", httpPort)
	return srv.ListenAndServe()
}

func addServices(ctx context.Context, grpcPort string) {
	optFuncs := []OptionFunc{
		eventCommand(),
		eventQuery(),
	}

	for _, optFunc := range optFuncs {
		optFunc(ctx, "localhost:"+grpcPort)
	}
}
