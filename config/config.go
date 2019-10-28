package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var confEnv *map[string]string
var envInit sync.Once

// GetConfig ..
func GetConfig() *map[string]string {
	envInit.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file: ", err.Error())
		}
		myEnv, err := godotenv.Read()
		if err != nil {
			log.Fatal("Error reading .env file: ", err.Error())
		}
		confEnv = &myEnv
	})
	return confEnv
}
