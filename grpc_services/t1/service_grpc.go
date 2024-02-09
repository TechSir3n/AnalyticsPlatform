package main

import (
	"context"
	hp "github.com/TechSir3n/analytics-platform/assistance"
	pb "github.com/TechSir3n/analytics-platform/grpc_services/t1/proto_buffer"
	"github.com/TechSir3n/analytics-platform/kafka/producer"
	"google.golang.org/grpc"
	"net"
	"os"
)

type GRPCServer struct {
	pb.UnimplementedOrderServiceServer
}

func newGRPCService() *GRPCServer {
	return &GRPCServer{}
}

func (s *GRPCServer) HandlerOrder(ctx context.Context, request *pb.OrderRequest) (*pb.OrderResponse, error) {
	if request.Id == "" && request.Name == "" && request.Type == "" && request.Amount < 0.0 && request.Time == "" {
		return &pb.OrderResponse{
			Status:      hp.Success,
			Description: "Incorrect data request",
		}, nil
	}

	apache := producer.OrderAndProduct{}
	apache.SetDataTrans(request.Id, request.Name, request.Type, request.Time, request.Amount)
	producer.SetObject(&apache)

	return &pb.OrderResponse{
		Status:      hp.Success,
		Description: "Order created successfuly",
	}, nil
}

func (s *GRPCServer) runGRPCService() error {
	conn, err := net.Listen(os.Getenv("GRPC_NETWORK"), os.Getenv("GRPC_ADDR"))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &GRPCServer{})

	if err := serv.Serve(conn); err != nil {
		return err
	}

	return nil
}
