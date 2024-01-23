package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	if err := runApacheKafka(); err != nil {
		panic(err)
	} else {
		log.Log.Info("apacheKafka success created")
	}
}
