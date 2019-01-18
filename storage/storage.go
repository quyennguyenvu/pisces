package storage

import (
	"context"
	"log"
	"pisces/config"
	"pisces/helper"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var db *mongo.Database
var once sync.Once

// Connect database
func Connect() {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		conf := config.GetConnection()
		client, err := mongo.Connect(ctx, conf.DataSource)
		if err != nil {
			helper.Logging("Storage", "Connect", err.Error())
		}
		err = client.Connect(ctx)
		if err != nil {
			helper.Logging("Storage", "Connect", err.Error())
		}
		log.Println("Database connected")
		db = client.Database(conf.Database)
	})
}
