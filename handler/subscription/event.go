package subscription

import (
	"log"
	"pisces/pb/v1/event"
	"strconv"

	nats "github.com/nats-io/go-nats"
)

// WithEventCommand ..
func subscribeStore() OptionFunc {
	return func(natsConn *nats.EncodedConn) {
		natsConn.QueueSubscribe("Event.Store", "Events", func(m *event.StoreResponse) {
			log.Println("new id is " + strconv.FormatUint(m.Id, 10))
		})
	}
}
