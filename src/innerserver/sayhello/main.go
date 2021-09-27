//sayhellow的服务端实现

package main

import (
	"context"
	pb "gm/common/in"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
