package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	grpc := newGRPCService()
	if err := grpc.runGRPCService(); err != nil {
		log.Log.Panic(err)
	} else {
		log.Log.Info("Success run the Product microservice")
	}
}
