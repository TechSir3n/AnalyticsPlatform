package main

import (
	"context"
	"net"
	"os"
	hp "github.com/TechSir3n/analytics-platform/assistance"
	pb "github.com/TechSir3n/analytics-platform/grpc_services/t1/proto_buffer"
	"github.com/TechSir3n/analytics-platform/kafka/producer"
	cmd "github.com/TechSir3n/analytics-platform/kafka/producer/cmd"
	"google.golang.org/grpc"
)

type GRPCServiceTransaction struct {
	pb.UnimplementedOrderServiceServer
}

func newGRPCServiceTransaction() *GRPCServiceTransaction {
	return &GRPCServiceTransaction{}
}

func (s *GRPCServiceTransaction) HandlerOrder(ctx context.Context, request *pb.OrderRequest) (*pb.OrderResponse, error) {
	if request.Id == "" && request.Name == "" && request.Type == "" && request.Amount < 0.0 && request.Time == "" {
		return &pb.OrderResponse{
			Status:      hp.Success,
			Description: "Incorrect data request",
		}, nil
	}

	producer.SendApacheBrokerTrans(request.Id, request.Name, request.Type, request.Time, request.Amount, cmd.Producer)

	return &pb.OrderResponse{
		Status:      hp.Success,
		Description: "Order created successfuly",
	}, nil
}

func (s *GRPCServiceTransaction) runGRPCService() error {
	conn, err := net.Listen(os.Getenv("GRPC_NETWORK"), os.Getenv("GRPC_ADDR_TRANSACTION"))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterOrderServiceServer(serv, &GRPCServiceTransaction{})

	if err := serv.Serve(conn); err != nil {
		return err
	}

	return nil
}
