package main

import (
	"context"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"k8s-go-grpc/proto"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	for *b != 0 {
		res := *a % *b
		a, b = b, &res
	}
	return &pb.GCDResponse{Result: a}, nil
}