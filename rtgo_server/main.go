package main

import (
	"log"
	"net"
	pb "rt1go/protos/rtgo/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRtgoServer
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRtgoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
