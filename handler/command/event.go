package command

import (
	"context"
	"pisces/handler"
	"pisces/pb/v1/event"
	"pisces/service"

	log "github.com/Sirupsen/logrus"
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
		log.WithFields(log.Fields{
			"entity": "Event command",
			"method": "Store",
		}).Warning(response.Err.Error())
		return nil, response.Err
	}

	storedEvent := &event.StoreResponse{
		Api:     handler.APIVersion,
		Id:      uint64(response.Data.(uint)),
		Success: true,
	}

	// publish message
	natsConn := handler.GetConnection()
	if err := natsConn.Publish(
		"Event.Store",
		storedEvent,
	); err != nil {
		log.WithFields(log.Fields{
			"entity": "Event command",
			"method": "Publish",
		}).Warning(err.Error())
	}

	return storedEvent, nil
}
