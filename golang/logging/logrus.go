package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// set json formatter
	log.SetFormatter(&log.JSONFormatter{})
	// set log level
	log.SetLevel(log.WarnLevel)

	log.WithFields(log.Fields{
		"user": "admin",
	}).Info("Some interesting info")

	log.Warn("This is a warning")
	log.Error("An error occured!")
}
