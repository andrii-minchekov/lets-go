package main

import (
	"fmt"
	"github.com/andrii-minchekov/lets-go/pkg/models"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockTestifyDb struct {
	mock.Mock
}

func (db *MockTestifyDb) FindSnippets() (models.Snippets, error) {
	fmt.Println("Mocked find snippets function")

	db.Called()
	return models.Snippets{&models.Snippet{ID: 1}}, nil
}

func (db *MockTestifyDb) InsertSnippet(title, content string) (int, error) {
	panic("implement me")
}

func (db *MockTestifyDb) LatestSnippets() (models.Snippets, error) {
	panic("implement me")
}

func (db *MockTestifyDb) GetSnippet(id int) (*models.Snippet, error) {
	panic("implement me")
}

func (db *MockTestifyDb) InsertUser(name, email, password string) error {
	panic("implement me")
}

func (db *MockTestifyDb) VerifyUser(email, password string) (int, error) {
	panic("implement me")
}

func TestSnippetsService_GetSnippets(t *testing.T) {
	mockDb := new(MockTestifyDb)

	mockDb.On("FindSnippets").Return(models.Snippets{&models.Snippet{ID: 1}})

	s := SnippetsService{mockDb}

	s.findAllSnippets()

	mockDb.AssertNumberOfCalls(t, "FindSnippets", 1)
	mockDb.AssertExpectations(t)
}
