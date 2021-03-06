package uc

import (
	"github.com/andrii-minchekov/lets-go/domain/snippet"
	"log"
)

type SnippetUseCase interface {
	CreateSnippet(snippet snp.Snippet) (int64, error)
	LatestSnippets() (snp.Snippets, error)
	GetSnippet(id int64) (*snp.Snippet, error)
}

type snippetUseCaseImpl struct {
	Repo snp.SnippetRepository
}

func NewSnippetUseCase(repo snp.SnippetRepository) SnippetUseCase {
	if repo == nil {
		log.Panicf("repo shouldn't be null")
	}
	return snippetUseCaseImpl{Repo: repo}
}

func (useCase snippetUseCaseImpl) CreateSnippet(snippet snp.Snippet) (int64, error) {
	id, err := useCase.Repo.AddSnippet(snippet)
	if err != nil {
		log.Printf("Couldn't add snippet %s", err)
	} else {
		log.Printf("CreateSnippet UseCase created snippet with id=%d", id)
	}
	return id, err
}

func (useCase snippetUseCaseImpl) LatestSnippets() (snp.Snippets, error) {
	snippets, err := useCase.Repo.LatestSnippets()
	if err != nil {
		log.Printf("Couldn't get latest snippets %s", err)
	}
	log.Printf("Latest snippets count is %d", len(snippets))
	return snippets, err
}

func (useCase snippetUseCaseImpl) GetSnippet(id int64) (*snp.Snippet, error) {
	snippet, err := useCase.Repo.GetSnippet(id)
	if err != nil {
		log.Printf("Couldn't get snippet by id %d because %s", id, err)
	} else {
		log.Printf("Snippet by id %d found successfully", id)
	}
	return snippet, err
}
