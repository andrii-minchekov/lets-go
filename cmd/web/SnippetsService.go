package main

import (
	"fmt"
	"github.com/andrii-minchekov/lets-go/pkg/models"
)

type SnippetsService struct {
	Database models.Db
}

func (s *SnippetsService) findAllSnippets() (models.Snippets, error) {
	snippets, err := s.Database.FindSnippets()

	if err != nil {
		fmt.Errorf("error was caught while executing a query")
	}

	if snippets == nil {
		fmt.Errorf("snippets is empty")
	}

	return snippets, nil
}
