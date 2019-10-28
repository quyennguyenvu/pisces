package service

import (
	"encoding/json"
	"net/http"
	"pisces/pb/v1/event"
	"pisces/storage"
)

// EventService ..
type EventService interface {
	Store(request *event.StoreRequest) Response
	// ByID(strID string) Response
	List(request *event.ListRequest) Response
}

// eventServiceImpl ..
type eventServiceImpl struct {
	eventDS storage.EventStorage
}

// NewEventService ..
func NewEventService() EventService {
	return &eventServiceImpl{
		eventDS: storage.NewEventStorage(),
	}
}

// Store ..
func (sc *eventServiceImpl) Store(request *event.StoreRequest) Response {
	event := storage.Event{
		Body: request.Body,
	}

	inserted, err := sc.eventDS.Store(&event)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return Response{
		Data: inserted.ID,
		Code: http.StatusOK,
		Err:  nil,
	}
}

// List ..
func (sc *eventServiceImpl) List(request *event.ListRequest) Response {
	events, err := sc.eventDS.List()
	if err != nil || events == nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	res, err := json.Marshal(events)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return Response{
		Data: res,
		Code: http.StatusOK,
		Err:  nil,
	}
}
