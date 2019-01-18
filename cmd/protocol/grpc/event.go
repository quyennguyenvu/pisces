package grpc

import (
	"pisces/pb/v1/event"

	"google.golang.org/grpc"
)

// WithEventCommand ..
func WithEventCommand(eventSrv event.EventCommandServer) OptionFunc {
	return func(srv *grpc.Server) {
		event.RegisterEventCommandServer(srv, eventSrv)
	}
}

// WithEventQuery ..
func WithEventQuery(eventSrv event.EventQueryServer) OptionFunc {
	return func(srv *grpc.Server) {
		event.RegisterEventQueryServer(srv, eventSrv)
	}
}
