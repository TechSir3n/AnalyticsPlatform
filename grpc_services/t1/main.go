package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
	cmd "github.com/TechSir3n/analytics-platform/kafka/producer/cmd"
)

func main() {
	serv := newGRPCServiceTransaction()
	if err := serv.runGRPCService(); err != nil {
		panic(err)
	} else {
		log.Log.Info("Success run the OrderMicroserivce")
	}

	defer cmd.CloseProducer()
}
