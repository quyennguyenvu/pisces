package storage

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

// Event ..
type Event struct {
	gorm.Model
	Body string
}

// TableName ...
func (m Event) TableName() string {
	return "events"
}

// EventStorage ..
type EventStorage interface {
	ByID(id int) (*Event, error)
	Store(data *Event) (*Event, error)
	List() (*[]*Event, error)
	Update(data *Event) error
	Destroy(data *Event) error
}

type eventImpl struct {
	db *gorm.DB
}

// NewEventStorage ..
func NewEventStorage() EventStorage {
	return &eventImpl{db: db}
}

func (s *eventImpl) Store(data *Event) (*Event, error) {
	result := s.db.Create(data)
	if result.Error != nil {
		log.WithFields(log.Fields{
			"entity": "Event",
			"method": "Store",
		}).Error(result.Error.Error())
		return nil, result.Error
	}
	return data, nil
}

func (s *eventImpl) List() (*[]*Event, error) {
	events := &[]*Event{}
	result := s.db.Find(events)
	if result.Error != nil {
		log.WithFields(log.Fields{
			"entity": "Event",
			"method": "List",
		}).Error(result.Error.Error())
		return nil, result.Error
	}
	return events, nil
}

func (s *eventImpl) ByID(id int) (*Event, error) {
	event := &Event{}
	result := s.db.First(&event, id)
	if result.Error != nil {
		log.WithFields(log.Fields{
			"entity": "Todo",
			"method": "ByID",
		}).Error(result.Error.Error())
		return nil, result.Error
	}
	return event, nil
}

func (s *eventImpl) Update(data *Event) error {
	result := s.db.Save(data)
	if result.Error != nil {
		log.WithFields(log.Fields{
			"entity": "Event",
			"method": "Update",
		}).Error(result.Error.Error())
		return result.Error
	}
	return nil
}

func (s *eventImpl) Destroy(data *Event) error {
	result := s.db.Delete(data)
	if result.Error != nil {
		log.WithFields(log.Fields{
			"entity": "Event",
			"method": "Destroy",
		}).Error(result.Error.Error())
		return result.Error
	}
	return nil
}
