package db

import (
	snp "github.com/andrii-minchekov/lets-go/domain/snippet"
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrmSnippetRepository_GetUserSnippetSuccessIT(t *testing.T) {
	CleanupDb()
	repo := OrmSnippetRepository{
		Db: GetDatabase(),
	}
	snippetId, _ := CreateUserAndSnippet()

	got, err := repo.GetUserSnippet(snippetId)

	if err != nil {
		t.Fatalf("GetUserSnippet() error = %v", err)
	}
	require.Equal(t, snippetId, got.Snippet.ID)
	require.Greater(t, got.User.Id, int64(0))
}

func CreateUserAndSnippet() (int64, error) {
	snpRepository := NewOrmSnippetRepository()
	usrRepository := NewDbUserRepository()
	userId, _ := usrRepository.CreateUser(usr.User{
		Name:     "nameCreateUserAndSnippet",
		Email:    "CreateUserAndSnippet@example.com",
		Password: "12",
	})
	return snpRepository.AddSnippet(snp.Snippet{
		Title:   "titleCreateUserAndSnippet",
		Content: "contentCreateUserAndSnippet",
		UserId:  &userId,
	})
}
