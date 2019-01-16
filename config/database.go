package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Connection ..
type Connection struct {
	DataSource string
	Database   string
}

// Serve ..
type Serve struct {
	Port string
}

var confConnection *Connection
var confServe *Serve

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	confConnection = &Connection{
		DataSource: dataSource,
		Database:   database,
	}

	servePort := os.Getenv("SERVE_PORT")
	confServe = &Serve{
		Port: servePort,
	}
}

// GetConnection ..
func GetConnection() *Connection {
	return confConnection
}

// GetServe ..
func GetServe() *Serve {
	return confServe
}
