package subscription

import (
	"log"
	"pisces/handler"

	nats "github.com/nats-io/go-nats"
)

// OptionFunc ..
type OptionFunc func(*nats.EncodedConn)

// Run ..
func Run() {
	log.Println("Subscribe bitches!!")

	natsConn := handler.GetConnection()
	optFuncs := []OptionFunc{
		subscribeStore(),
	}

	for _, optFunc := range optFuncs {
		optFunc(natsConn)
	}
}
