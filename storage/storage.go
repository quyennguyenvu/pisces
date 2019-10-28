package storage

import (
	"pisces/config"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var once sync.Once

// Connect database
func Connect() {
	once.Do(func() {
		conf := config.GetConnection()
		var err error
		db, err = gorm.Open(conf.Driver, conf.DataSource)
		if err != nil {
			log.WithFields(log.Fields{
				"entity": "Storage",
				"method": "Connect",
			}).Error(err.Error())
		}
		log.Println("Database connected")
		db.LogMode(conf.LogMode)
	})
}

// Disconnect ..
func Disconnect() {
	db.Close()
}
