package command

import (
	"context"
	"log"
	"pisces/handler"
	"pisces/helper"
	"pisces/pb/v1/event"
	"pisces/service"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// EventCommand ..
type EventCommand interface {
	Store(ctx context.Context, in *event.StoreRequest) (*event.StoreResponse, error)
}

type eventCommandImpl struct {
	eventSC service.EventService
}

// NewEventCommand ..
func NewEventCommand() EventCommand {
	return &eventCommandImpl{
		eventSC: service.NewEventService(),
	}
}

// Store ..
func (h *eventCommandImpl) Store(ctx context.Context, request *event.StoreRequest) (*event.StoreResponse, error) {
	log.Println("Creating event")
	response := h.eventSC.Store(request)
	if response.Err != nil {
		helper.Logging("Event", "Store", response.Err.Error())
		return nil, response.Err
	}

	storedEvent := &event.StoreResponse{
		Api:     handler.APIVersion,
		Id:      response.Data.(primitive.ObjectID).Hex(),
		Success: true,
	}

	// publish message
	natsConn := handler.GetConnection()
	if err := natsConn.Publish(
		"Event.Store",
		storedEvent,
	); err != nil {
		helper.Logging("Event", "Store", err.Error())
	}

	return storedEvent, nil
}
