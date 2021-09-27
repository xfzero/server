package main

import (
	"context"
	"fmt"
	pb "gm/common/in"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	RET_OK         = 1
	RET_PARAMS_ERR = 1001
)

type grpcService struct{}

func (s *grpcService) GMCmd(ctx context.Context, in *pb.GMRequest) (*pb.GMReplay, error) {
	req := fmt.Sprintf("%+v", in)
	fmt.Println("GMRequest:", req)

	//处理逻辑done：调用方法，方法中通过cmd来处理不同的逻辑

	replay := &pb.GMReplay{
		Ret: RET_OK,
		Cmd: in.GetCmd(),
		Sid: in.GetSid(),
	}

	return replay, nil
}

func (s *grpcService) GRPCTest(ctx context.Context, in *pb.TestRequest) (*pb.TestReplay, error) {
	return &pb.TestReplay{Replay: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWebServiceServer(s, &grpcService{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
