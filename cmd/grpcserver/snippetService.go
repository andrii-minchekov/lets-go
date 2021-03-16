package main

import (
	"context"
	pb "github.com/andrii-minchekov/lets-go/app/impl/grpc"
	uc "github.com/andrii-minchekov/lets-go/app/usecases"
	snp "github.com/andrii-minchekov/lets-go/domain/snippet"
)

type snippetServer struct {
	useCases uc.UseCases
}

func (s *snippetServer) CreateSnippet(c context.Context, r *pb.CreateSnippetRequest) (*pb.CreateSnippetResponse, error) {
	id, err := s.useCases.CreateSnippet(snp.Snippet{
		Title:   r.Title,
		Content: r.Content,
	})
	if err != nil {
		return nil, err
	}
	response := &pb.CreateSnippetResponse{Id: id}
	return response, nil
}
