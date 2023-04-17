package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/server"
	"net"
)

type Server struct {
	pb.UnimplementedSayHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + in.RequestName}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":9090")

	if err != nil {
		panic("network panic")
	}

	gserver := grpc.NewServer()

	pb.RegisterSayHelloServer(gserver, &Server{})

	gserver.Serve(listen)
}
