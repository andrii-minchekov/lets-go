package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/andrii-minchekov/lets-go/app/impl/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewSnippetServiceClient(conn)
	rand.Seed(time.Now().UnixNano())
	title := flag.String("title", "title"+strconv.Itoa(rand.Intn(1000)), "Specify title")
	content := flag.String("content", "content"+strconv.Itoa(rand.Intn(1000)), "Specify content")
	flag.Parse()
	request := &pb.CreateSnippetRequest{
		Title:   *title,
		Content: *content,
	}
	response, err := client.CreateSnippet(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Printf("Response: %d", response.Id)
}
