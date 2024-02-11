package main

import (
	cmd "github.com/TechSir3n/analytics-platform/kafka/producer/cmd"
	log "github.com/TechSir3n/analytics-platform/logging"
)

func main() {
	serv := newGRPCServiceProduct()
	if err := serv.runGRPCService(); err != nil {
		log.Log.Panic(err)
	} else {
		log.Log.Info("Success run the Product microservice")
	}

	defer cmd.CloseProducer()
}
