package main

import (
	"context"
	pb "grpc-test/grpctest"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *server) GetFeature(ctx context.Context, in *pb.Person) (*pb.Feature, error) {
	return &pb.Feature{Name: "Alice"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})

	// grep_cli 用
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
