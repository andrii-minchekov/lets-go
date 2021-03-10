package main

import (
	"fmt"
	"github.com/andrii-minchekov/lets-go/app/impl"
	pb "github.com/andrii-minchekov/lets-go/app/impl/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	port := "9090"
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("gRPC server is listening on port %s", port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	useCases := impl.NewComposedUseCases(impl.NewFlagConfig())
	pb.RegisterSnippetServiceServer(grpcServer, &snippetServer{useCases})
	grpcServer.Serve(listener)
}
