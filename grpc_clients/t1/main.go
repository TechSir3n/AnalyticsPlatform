package main

import (
	"log"
)

func main() {
	grpc := newGRPClient()
	if err := grpc.runClientGRPC(); err != nil {
		log.Panic(err)
	}

	log.Println("Successed run grpClient")
}
