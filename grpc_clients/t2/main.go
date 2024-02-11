package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	grpc := newGRPClientProduct()
	if err := grpc.grpClientRun(); err != nil {
		log.Log.Panic(err)
	} else {
		log.Log.Info("Success Run grpClientRun [Product]")
	}
}
