package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mannanmcc/transferMoney/proto"
	"google.golang.org/grpc"
)

type Server struct{}

//TransferMoney implement the TransferMoney
func (s *Server) TransferMoney(ctx context.Context, res *proto.TransferMoneyRequest) (*proto.TransferMoneyResponse, error) {
	result := &proto.TransferMoneyResponse{
		Rate: 2123.23,
	}

	return result, nil
}

func main() {
	fmt.Println("Transfering money for the clinet")

	lis, err := net.Listen("tcp", "0.0.0.0:5002")
	if err != nil {
		log.Fatalf("oops failed to listen on port: 5002 :%v", err)
	}

	s := grpc.NewServer()
	proto.RegisterTransferServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error :%v", err)
	}
}
