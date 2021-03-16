package uc

import (
	"errors"
	"github.com/andrii-minchekov/lets-go/app/impl/db"
	snp "github.com/andrii-minchekov/lets-go/domain/snippet"
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"github.com/stretchr/testify/require"
	"log"
	"strconv"
	"testing"
)

type mockRepo struct {
	expectedErr  error
	expSnippetId int64
}

func (r mockRepo) GetUserSnippet(id int64) (*snp.UserSnippet, error) {
	panic("implement me")
}

func (r mockRepo) LatestSnippets() (snp.Snippets, error) {
	panic("implement me")
}

func (r mockRepo) GetSnippet(id int64) (*snp.Snippet, error) {
	panic("implement me")
}

func (r mockRepo) AddSnippet(snippet snp.Snippet) (int64, error) {
	log.Printf("Stub AddSnippet is called")
	if r.expectedErr != nil {
		return 0, r.expectedErr
	}
	return r.expSnippetId, nil
}

func TestSnippetUseCase_CreateSnippetUT(t *testing.T) {
	type fields struct {
		Repo snp.SnippetRepository
	}
	type args struct {
		snippet snp.Snippet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			"should create snippet successfully when snippet input is valid",
			fields{Repo: mockRepo{expSnippetId: int64(1)}},
			args{snp.Snippet{}},
			1,
			false,
		},
		{
			"should return expectedErr when repo fails to add snippet",
			fields{Repo: mockRepo{expectedErr: errors.New("couldn't add snippet")}},
			args{snp.Snippet{}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewSnippetUseCase(tt.fields.Repo)
			got, err := useCase.CreateSnippet(tt.args.snippet)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSnippet() expectedErr = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSnippet() got = %v, wantValue %v", got, tt.want)
			}
		})
	}
}

func TestSnippetUseCase_LatestSnippetsIT(t *testing.T) {
	type fields struct {
		Repo snp.SnippetRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "should return latest created snippets",
			fields:  fields{db.NewOrmSnippetRepository()},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := NewSnippetUseCase(tt.fields.Repo)
			for i := 0; i < 11; i++ {
				useCase.CreateSnippet(snp.Snippet{
					Title:   "title" + strconv.Itoa(i),
					Content: "content" + strconv.Itoa(i),
				})
			}
			got, err := useCase.LatestSnippets()
			if (err != nil) != tt.wantErr {
				t.Errorf("LatestSnippets() expectedErr = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("LatestSnippets() got = %v, wantValue %v", len(got), tt.want)
			}
		})
	}
}

func TestGetSnippetSuccessIT(t *testing.T) {
	snpUseCase, snippetId := CreateUserAndSnippet()

	actual, err := snpUseCase.GetSnippet(int64(snippetId))

	require.NotEmptyf(t, actual, "Empty because of %v", err)
	require.NotEmptyf(t, actual.UserId, "userId of snippet should not be empty")
	require.Greater(t, actual.ID, int64(0))
	require.EqualValues(t, "TitleTestGetSnippetSuccessIT", actual.Title)
}

func CreateUserAndSnippet() (SnippetUseCase, SnippetId) {
	snpUseCase := NewSnippetUseCase(db.NewOrmSnippetRepository())
	usrUseCase := NewUserUseCase(db.NewDbUserRepository())
	userId, _ := usrUseCase.SignupUser(usr.User{
		Name:     "andi",
		Email:    "andi@example.com",
		Password: "12",
	})
	inSnippet := snp.Snippet{
		Title:   "TitleTestGetSnippetSuccessIT",
		Content: "ContentTestGetSnippetSuccessIT",
		UserId:  &userId,
	}
	snippetId, _ := snpUseCase.CreateSnippet(inSnippet)
	return snpUseCase, SnippetId(snippetId)
}

type SnippetId int64
