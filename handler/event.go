package handler

import (
	"context"
	"encoding/json"
	"gemini/pb/v1/event"
	"gemini/service"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// EventHandler ..
type EventHandler interface {
	Store(ctx context.Context, in *event.StoreRequest) (*event.StoreResponse, error)
	List(ctx context.Context, filter *event.ListRequest) (*event.ListResponse, error)
}

type eventHandlerImpl struct {
	eventSC service.EventService
}

// NewEventHandler ..
func NewEventHandler() EventHandler {
	return &eventHandlerImpl{
		eventSC: service.NewEventService(),
	}
}

// Store ..
func (h *eventHandlerImpl) Store(ctx context.Context, request *event.StoreRequest) (*event.StoreResponse, error) {
	log.Println("Creating event")
	response := h.eventSC.Store(request)
	if response.Err != nil {
		return nil, response.Err
	}
	return &event.StoreResponse{
		Api:     apiVersion,
		Id:      response.Data.(primitive.ObjectID).Hex(),
		Success: true,
	}, nil
}

// List ..
func (h *eventHandlerImpl) List(ctx context.Context, filter *event.ListRequest) (*event.ListResponse, error) {
	log.Println("List events")
	response := h.eventSC.List()
	if response.Err != nil {
		return nil, response.Err
	}

	var data []*event.Event
	if err := json.Unmarshal(response.Data.([]byte), &data); err != nil {
		return nil, err
	}
	return &event.ListResponse{
		Api:    apiVersion,
		Events: data,
	}, nil
}
