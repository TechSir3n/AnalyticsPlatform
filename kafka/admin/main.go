package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	if err := createAdminApacheKafka(); err != nil {
		log.Log.Panic(err)
	} else {
		log.Log.Info("Kafka-Admin successed Run")
	}
}
