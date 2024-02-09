package main

import (
	"context"
	hp "github.com/TechSir3n/analytics-platform/assistance"
	pb "github.com/TechSir3n/analytics-platform/grpc_services/t2/proto_buffer"
	"github.com/TechSir3n/analytics-platform/kafka/producer"
	"google.golang.org/grpc"
	"net"
	"os"
)

type GRPCService struct {
	pb.UnimplementedProductServiceServer
}

func newGRPCService() *GRPCService {
	return &GRPCService{}
}

func (*GRPCService) CreaterProduct(ctx context.Context, request *pb.ProductRequest) (*pb.ProductResponse, error) {
	if request.Id == "" && request.Name == "" && request.Price < 0.0 && request.Quantity <= 0 {
		return &pb.ProductResponse{
			Status:      hp.Incorrect,
			Description: "Request data is incorrect [createProduct]",
		}, nil
	}

	apache := producer.OrderAndProduct{}
	apache.SetDataProduct(request.Id, request.Name, float64(request.Price), request.Quantity)
	producer.SetObject(&apache)

	return &pb.ProductResponse{
		Status:      hp.Success,
		Description: "Success created request [createProduct]",
	}, nil
}

func (g *GRPCService) runGRPCService() error {
	conn, err := net.Listen(os.Getenv("GRPC_NETWORK"), os.Getenv("GRPC_ADDR"))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterProductServiceServer(serv, &GRPCService{})

	if err := serv.Serve(conn); err != nil {
		return err
	}

	return nil
}
