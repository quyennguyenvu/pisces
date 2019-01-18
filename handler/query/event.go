package query

import (
	"context"
	"encoding/json"
	"log"
	"pisces/handler"
	"pisces/pb/v1/event"
	"pisces/service"
)

// EventQuery ..
type EventQuery interface {
	List(ctx context.Context, request *event.ListRequest) (*event.ListResponse, error)
}

type eventQueryImpl struct {
	eventSC service.EventService
}

// NewEventQuery ..
func NewEventQuery() EventQuery {
	return &eventQueryImpl{
		eventSC: service.NewEventService(),
	}
}

// List ..
func (h *eventQueryImpl) List(ctx context.Context, request *event.ListRequest) (*event.ListResponse, error) {
	log.Println("List events")
	response := h.eventSC.List(request)
	if response.Err != nil {
		return nil, response.Err
	}

	var data []*event.Event
	if err := json.Unmarshal(response.Data.([]byte), &data); err != nil {
		return nil, err
	}
	return &event.ListResponse{
		Api:    handler.APIVersion,
		Events: data,
	}, nil
}
