package storage

import (
	"context"
	"gemini/helper"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Event ..
type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Body      string             `bson:"body"`
	CreatedAt time.Time          `bson:"created_at"`
}

// EventStorage ..
type EventStorage interface {
	Store(data *Event) (interface{}, error)
	// ByID(id int) (*Event, error)
	List() ([]Event, error)
}

type eventImpl struct {
	coll *mongo.Collection
}

// NewEventStorage ..
func NewEventStorage() EventStorage {
	return &eventImpl{coll: db.Collection("events")}
}

func (s *eventImpl) Store(data *Event) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserted, err := s.coll.InsertOne(ctx, data)
	if err != nil {
		helper.Logging("Event", "Store", err.Error())
		return nil, err
	}
	return inserted.InsertedID, nil
}

func (s *eventImpl) List() ([]Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := s.coll.Find(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	events := []Event{}
	for cur.Next(ctx) {
		var event Event
		if err := cur.Decode(&event); err != nil {
			helper.Logging("Event", "List", err.Error())
			return nil, err
		}

		events = append(events, event)
	}
	if err := cur.Err(); err != nil {
		helper.Logging("Event", "List", err.Error())
		return nil, err
	}

	return events, nil
}
