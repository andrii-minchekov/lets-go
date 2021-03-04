package main

import (
	_ "fmt"
	"github.com/andrii-minchekov/lets-go/pkg/models"
	"testing"
)

type MockDb struct{}

func (db *MockDb) FindSnippets() (models.Snippets, error) {
	return models.Snippets{&models.Snippet{ID: 1}}, nil
}

func (db *MockDb) InsertSnippet(title, content string) (int, error) {
	return 1, nil
}

func (db *MockDb) LatestSnippets() (models.Snippets, error) {
	return models.Snippets{&models.Snippet{ID: 1}}, nil
}

func (db *MockDb) GetSnippet(id int) (*models.Snippet, error) {
	return &models.Snippet{ID: 1}, nil
}

func (db *MockDb) InsertUser(name, email, password string) error {
	return nil
}

func (db *MockDb) VerifyUser(email, password string) (int, error) {
	return 1, nil
}

func TestApp_GetSnippets(t *testing.T) {
	app := App{Database: &MockDb{}}
	snippets, err := app.Database.FindSnippets()

	if err != nil {
		t.Errorf("Test has failed with error")
	}

	for _, snippet := range snippets {
		if snippet.ID != 1 {
			t.Errorf("got %d, want %d", 1, snippet.ID)
		}
	}
	//req, _ := http.NewRequest("GET", "/snippets", nil)
	//w := httptest.NewRecorder()
	//homeHandle.ServeHTTP(w, req)
}
