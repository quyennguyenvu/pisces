package helper

import (
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

// Logging ..
func Logging(entity, method, msg string) {
	currentTime := time.Now()
	strCurr := currentTime.Format("2006-01-02")

	filename := "./logging/logfile-" + strCurr + ".log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer f.Close()

	if err != nil {
		log.Warn(err)
	} else {
		log.SetOutput(f)
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.WarnLevel)
	}

	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	log.WithFields(log.Fields{
		"entity": entity,
		"method": method,
	}).Info(msg)
}
