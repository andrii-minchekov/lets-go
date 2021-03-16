package db

import (
	"context"
	m "github.com/andrii-minchekov/lets-go/app/impl/db/models"
	"github.com/andrii-minchekov/lets-go/domain/snippet"
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"time"
)

type OrmSnippetRepository struct {
	Db Database
}

func NewOrmSnippetRepository() snp.SnippetRepository {
	return OrmSnippetRepository{Db: GetDatabase()}
}

func (r OrmSnippetRepository) AddSnippet(snippet snp.Snippet) (int64, error) {
	s := m.Snippet{
		Title:   snippet.Title,
		Content: snippet.Content,
		Created: time.Now().UTC(),
		UserID:  null.Int64FromPtr(snippet.UserId),
	}
	err := s.Insert(context.Background(), r.Db, boil.Infer())
	if err != nil {
		return 0, err
	}
	log.Printf("Snippet with id=%d successfully added through out orm", s.ID)
	return s.ID, nil
}

func (r OrmSnippetRepository) LatestSnippets() (snp.Snippets, error) {
	var slice []m.Snippet
	err := m.Snippets(qm.Limit(10)).Bind(context.Background(), r.Db, &slice)
	if err != nil {
		return nil, err
	}
	var result snp.Snippets
	for _, snippet := range slice {
		result = append(result, convToSnippet(&snippet))
	}
	log.Printf("Returned latest %d snippets", len(result))
	return result, nil
}

func (r OrmSnippetRepository) GetSnippet(id int64) (*snp.Snippet, error) {
	entity, err := m.Snippets(m.SnippetWhere.ID.EQ(id)).One(context.Background(), r.Db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return convToSnippet(entity), nil
}

func (r OrmSnippetRepository) GetUserSnippet(id int64) (*snp.UserSnippet, error) {
	entity := SnippetAndUser{}
	if err := m.Snippets(
		qm.Select(m.TableNames.Snippets+"."+m.SnippetColumns.ID, m.SnippetColumns.Title, m.SnippetColumns.Content, m.TableNames.Snippets+"."+m.SnippetColumns.Created, m.TableNames.Users+"."+m.UserColumns.ID, m.UserColumns.Name, m.UserColumns.Email, m.TableNames.Users+"."+m.UserColumns.Created),
		m.SnippetWhere.ID.EQ(id),
		qm.InnerJoin(m.TableNames.Users+" on "+m.TableNames.Snippets+"."+m.SnippetColumns.UserID+"="+m.TableNames.Users+"."+m.UserColumns.ID),
	).Bind(context.Background(), r.Db, &entity); err != nil {
		log.Print(err)
		return nil, err
	}
	log.Printf("Repository GetUserSnippet found snippet entity with id=%d", entity.Snippet.ID)
	return convToUserSnippet(entity), nil
}

func convToSnippet(record *m.Snippet) *snp.Snippet {
	return &snp.Snippet{
		ID:      record.ID,
		Title:   record.Title,
		Content: record.Content,
		Created: record.Created,
		Expires: record.Expires.Ptr(),
		UserId:  record.UserID.Ptr(),
	}
}

func convToUserSnippet(entity SnippetAndUser) *snp.UserSnippet {
	return &snp.UserSnippet{
		Snippet: snp.Snippet{
			ID:      entity.Snippet.ID,
			Title:   entity.Snippet.Title,
			Content: entity.Snippet.Content,
			Created: entity.Snippet.Created,
			Expires: entity.Snippet.Expires.Ptr(),
			UserId:  nil,
		},
		User: usr.User{
			Id:       entity.User.ID,
			Name:     entity.User.Name,
			Email:    entity.User.Email,
			Password: entity.User.Password,
		},
	}
}

type SnippetAndUser struct {
	Snippet m.Snippet `boil:"snippets,bind"`
	User    m.User    `boil:"users,bind"`
}
