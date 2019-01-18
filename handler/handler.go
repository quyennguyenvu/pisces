package handler

import (
	"context"
	"log"
	"pisces/config"
	"pisces/helper"
	"sync"
	"time"

	nats "github.com/nats-io/go-nats"
	natsp "github.com/nats-io/go-nats/encoders/protobuf"
)

const (
	// APIVersion is version of API is provided by server
	APIVersion = "v1"
)

var encodedConn *nats.EncodedConn
var once sync.Once

// Connect ..
func Connect() {
	once.Do(func() {
		_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		conf := config.GetNatsConnection()
		conn, err := nats.Connect(conf.URL)
		if err != nil {
			log.Fatal(err)
			helper.Logging("Handler", "Connect", err.Error())
		}
		encodedConn, err = nats.NewEncodedConn(conn, natsp.PROTOBUF_ENCODER)
		if err != nil {
			log.Fatal(err)
			helper.Logging("Handler", "Encoded Connect", err.Error())
		}
		log.Println("NATS connected")
	})
}

// Disconnect ..
func Disconnect() {
	encodedConn.Close()
}

// GetConnection ..
func GetConnection() *nats.EncodedConn {
	return encodedConn
}
