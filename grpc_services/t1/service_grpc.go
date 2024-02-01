package main

import (
	"context"
	pb "github.com/TechSir3n/analytics-platform/grpc_services/t1/proto_buffer"
	"github.com/TechSir3n/analytics-platform/kafka/producer"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	pb.UnimplementedOrderServiceServer
}

func newGRPCService() *GRPCServer {
	return &GRPCServer{}
}

func (s *GRPCServer) CreaterOrder(ctx context.Context, request *pb.OrderRequest) (*pb.OrderResponse, error) {
	if request == nil {
		return &pb.OrderResponse{
			Status:      "ERROR",
			Description: "Incorrect data request",
		}, nil
	}

	order := producer.NewOrderTransaction()
	producer.SetOrderTransaction(order)
	order.SetApacheKafka(request.Id, request.Name, request.Type, request.Amount)

	return &pb.OrderResponse{
		Status:      "Success",
		Description: "Order created successfuly",
	}, nil
}

func (s *GRPCServer) runGRPCService() error {
	conn, err := net.Listen("tcp", ":8080")
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
