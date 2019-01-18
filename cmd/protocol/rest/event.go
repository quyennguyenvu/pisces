package rest

import (
	"context"
	"pisces/pb/v1/event"
)

func eventCommand() OptionFunc {
	return func(ctx context.Context, port string) {
		event.RegisterEventCommandHandlerFromEndpoint(ctx, mux, port, dialOpts)
	}
}

func eventQuery() OptionFunc {
	return func(ctx context.Context, port string) {
		event.RegisterEventQueryHandlerFromEndpoint(ctx, mux, port, dialOpts)
	}
}
