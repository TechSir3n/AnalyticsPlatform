package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	grpc := newGRPClient()
	if status, description, err := grpc.runClientGRPC(); err != nil {
		log.Log.Panic(err)
	} else {
		log.Log.Println(status, " - ", description)
	}
}
