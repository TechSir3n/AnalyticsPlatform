package main

import (
	log "github.com/TechSir3n/analytics-platform/logging"
	"fmt"
)

func main() {
	grpc := newGRPClientTransaction()
	if err := grpc.runClientGRPC(); err != nil {
		log.Log.Panic(err)
	} 
	fmt.Println("Success run client grpc[Transaction]")
}
