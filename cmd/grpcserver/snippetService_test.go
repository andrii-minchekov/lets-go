package main

import (
	"context"
	pb "github.com/andrii-minchekov/lets-go/app/impl/grpc"
	uc "github.com/andrii-minchekov/lets-go/app/usecases"
	"github.com/andrii-minchekov/lets-go/app/usecases/mocks"
	snp "github.com/andrii-minchekov/lets-go/domain/snippet"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

const title = "title"
const content = "content"
const snippetId = 1

func Test_snippetServer_CreateSnippet(t *testing.T) {
	type fields struct {
		useCases uc.UseCases
	}
	type args struct {
		c context.Context
		r *pb.CreateSnippetRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateSnippetResponse
		wantErr bool
	}{
		{
			name:   "should create snippet successfully",
			fields: fields{useCases: mockUseCases()},
			args: args{context.Background(), &pb.CreateSnippetRequest{
				Title:   title,
				Content: content,
			}},
			want:    &pb.CreateSnippetResponse{Id: int64(snippetId)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &snippetServer{
				useCases: tt.fields.useCases,
			}
			got, err := s.CreateSnippet(tt.args.c, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSnippet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSnippet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockUseCases() *mocks.UseCases {
	cases := &mocks.UseCases{}
	cases.On("CreateSnippet", mock.MatchedBy(func(snippet snp.Snippet) bool {
		return snippet.Content == content && snippet.Title == title
	})).Return(int64(1), nil)
	return cases
}
