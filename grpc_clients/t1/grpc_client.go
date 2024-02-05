package main

import (
	"context"
	"os"
	"time"

	"github.com/TechSir3n/analytics-platform/assistance"
	pb "github.com/TechSir3n/analytics-platform/grpc_services/t1/proto_buffer"
	log "github.com/TechSir3n/analytics-platform/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPClient struct {
	Status      string
	Description string
}

func newGRPClient() *GRPClient {
	return &GRPClient{}
}

func (c *GRPClient) runClientGRPC() error {
	conn, err := grpc.Dial(os.Getenv("GRPC_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		transaction := assistance.GenerateTransaction()

		client := pb.NewOrderServiceClient(conn)	
		res, err := client.HandlerOrder(context.Background(),
			&pb.OrderRequest{Id: transaction.ID, Name: transaction.Name, Type: transaction.Type,
				Time: transaction.Date.String(), Amount: transaction.Amount})
		if err != nil {
			return err
		}

		log.Log.Println(res.Description, res.Status)
		time.Sleep(time.Second * 2)
	}

	return nil
}
