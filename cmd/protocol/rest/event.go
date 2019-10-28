package rest

import (
	"context"
	"pisces/pb/v1/event"
)

// WithEventCommand ..
func WithEventCommand() OptionFunc {
	return func(ctx context.Context, port string) {
		event.RegisterEventCommandHandlerFromEndpoint(ctx, mux, port, dialOpts)
	}
}

// WithEventQuery ..
func WithEventQuery() OptionFunc {
	return func(ctx context.Context, port string) {
		event.RegisterEventQueryHandlerFromEndpoint(ctx, mux, port, dialOpts)
	}
}
