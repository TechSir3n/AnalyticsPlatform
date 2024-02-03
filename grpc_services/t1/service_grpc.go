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

func (s *GRPCServer) HandlerOrder(ctx context.Context, request *pb.OrderRequest) (*pb.OrderResponse, error) {
	if request.Id == "" && request.Name == "" && request.Type == "" && request.Amount < 0.0 && request.Time == "" {
		return &pb.OrderResponse{
			Status:      "ERROR",
			Description: "Incorrect data request",
		}, nil
	}

	order := producer.NewOrderTransaction()
	producer.SetOrderTransaction(order)
	order.SetApacheKafka(request.Id, request.Name, request.Type, request.Time, request.Amount)
	order.ApacheKafkaProducerRun()

	return &pb.OrderResponse{
		Status:      "Success",
		Description: "Order created successfuly",
	}, nil
}

func (s *GRPCServer) runGRPCService() error {
	conn, err := net.Listen("tcp", ":8010")
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
