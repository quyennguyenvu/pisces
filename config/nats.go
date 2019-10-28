package config

import (
	"sync"
)

// NatsConnection ..
type NatsConnection struct {
	URL string
}

var confNats *NatsConnection
var natsInit sync.Once

// GetNatsConnection ..
func GetNatsConnection() *NatsConnection {
	natsInit.Do(func() {
		confNats = &NatsConnection{
			URL: (*confEnv)["NATS_URL"],
		}
	})
	return confNats
}
