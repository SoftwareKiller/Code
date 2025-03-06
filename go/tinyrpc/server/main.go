package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"tinyrpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Hello(ctx context.Context, req *pb.String) (*pb.String, error) {
	return &pb.String{Value: fmt.Sprintf("%s, Hello -- %s",
		time.Now().String(), req.GetValue())}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalf("Failed to listen :%s", err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%s", err.Error())
	}
}
