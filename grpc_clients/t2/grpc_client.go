package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/TechSir3n/analytics-platform/assistance"
	pb "github.com/TechSir3n/analytics-platform/grpc_services/t2/proto_buffer"
	log "github.com/TechSir3n/analytics-platform/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPClientProduct struct {
	Status      string
	Description string
}

func newGRPClientProduct() *GRPClientProduct {
	return &GRPClientProduct{}
}

func (g *GRPClientProduct) grpClientRun() error {
	conn, err := grpc.Dial(os.Getenv("GRPC_ADDR_PRODUCT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	defer conn.Close()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	client := pb.NewProductServiceClient(conn)

	for range ticker.C {
		product := assistance.GenerateProduct()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		res, err := client.CreaterProduct(ctx, &pb.ProductRequest{Id: product.ID, Name: product.Name,
			Price: float32(product.Price), Quantity: int64(product.Quantity)})
		if err != nil {
			log.Log.Error("Error while calling CreateProduct: ", err)
			continue
		}

		fmt.Println(res.Status, res.Description)
	}

	return nil
}
