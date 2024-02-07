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

	client := pb.NewOrderServiceClient(conn)

	for range ticker.C {
		transaction := assistance.GenerateTransaction()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		res, err := client.HandlerOrder(ctx,
			&pb.OrderRequest{Id: transaction.ID, Name: transaction.Name, Type: transaction.Type,
				Time: transaction.Date.String(), Amount: transaction.Amount})
		if err != nil {
			log.Log.Error("Error while calling HandlerOrder: ", err)
			continue
		}

		log.Log.Println(res.Description, res.Status)
		time.Sleep(time.Second * 4)
	}

	return nil
}
