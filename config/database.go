package config

import (
	"os"
	"sync"
)

// DBConnection ..
type DBConnection struct {
	DataSource string
	Database   string
}

var confDB *DBConnection
var dbInit sync.Once

// GetConnection ..
func GetConnection() *DBConnection {
	dbInit.Do(func() {
		drive := os.Getenv("DB_CONNECTION")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		database := os.Getenv("DB_DATABASE")
		userName := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")

		// select datasource
		dataSource := drive + "://" +
			userName + ":" + password +
			"@" + dbHost + ":" + dbPort

		confDB = &DBConnection{
			DataSource: dataSource,
			Database:   database,
		}
	})
	return confDB
}
