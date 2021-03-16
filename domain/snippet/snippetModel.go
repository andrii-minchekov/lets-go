package snp

import (
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"time"
)

type Snippet struct {
	ID      int64
	Title   string
	Content string
	Created time.Time
	Expires *time.Time
	UserId  *int64
}

type UserSnippet struct {
	Snippet Snippet
	User    usr.User
}

type Snippets []*Snippet
