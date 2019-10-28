package helper

import (
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

var f *os.File

// Logging ..
func Logging() {
	currentTime := time.Now()
	strCurr := currentTime.Format("2006-01-02")

	filename := "./logging/logfile-" + strCurr + ".log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)

	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

// CloseFile ..
func CloseFile() {
	f.Close()
}
