package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
	"os"
)

func main() {
	if err := runApacheKafka(); err != nil {
		log.Log.Panic(err)
	} else {
		log.Log.Info("Kafka-Admin successed Run")
	}
	
	log.Log.Out.(*os.File).Close()
}
