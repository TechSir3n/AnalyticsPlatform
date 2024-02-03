package main

import (
	"context"
	"fmt"
	"time"

	"github.com/TechSir3n/analytics-platform/assistance"
	pb "github.com/TechSir3n/analytics-platform/grpc_services/t1/proto_buffer"
	"google.golang.org/grpc"
)

type GRPClient struct {
	Status      string
	Description string
}

func newGRPClient() *GRPClient {
	return &GRPClient{}
}

func (c *GRPClient) runClientGRPC() (string, string, error) {
	conn, err := grpc.Dial(":8010", grpc.WithInsecure())
	if err != nil {
		return "", "", err
	}
	defer conn.Close()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		transaction := assistance.GenerateTransaction()

		client := pb.NewOrderServiceClient(conn)
		res, err := client.HandlerOrder(context.Background(),
			&pb.OrderRequest{Id: transaction.ID, Name: transaction.Name,
				Type: transaction.Type, Time: transaction.Date.String(), Amount: transaction.Amount})
		if err != nil {
			return "", "", err
		}

		fmt.Println(res.Description, res.Status)
		time.Sleep(time.Second * 2)
	}

	return "", "", nil
}
