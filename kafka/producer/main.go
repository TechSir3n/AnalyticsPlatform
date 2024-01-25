package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	if err := apacheKafkaProducer(); err != nil {
		log.Log.Panic("Failed to run Apache-Kafka producer: " + err.Error())
	} else {
		log.Log.Info("Apache-Kafka Producer successed run")
	}
}
