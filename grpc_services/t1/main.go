package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	serv := newGRPCService()
	if err := serv.runGRPCService(); err != nil {
		panic(err)
	} else {
		log.Log.Info("Success run the first microservice")
	}
}
