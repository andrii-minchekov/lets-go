package db

import (
	"database/sql"
	"errors"
	"github.com/andrii-minchekov/lets-go/domain/snippet"
)

type DbSnippetRepository struct {
	Db Database
}

func NewDbSnippetRepository() DbSnippetRepository {
	return DbSnippetRepository{GetDatabase()}
}

func (r DbSnippetRepository) AddSnippet(snippet snp.Snippet) (int64, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
	VALUES($1, $2, CURRENT_DATE, CURRENT_DATE + INTERVAL '100000 seconds') returning id`

	result, err := r.Db.Query(stmt, snippet.Title, snippet.Content)

	if err != nil {
		return 0, err
	}

	if err := result.Next(); !err {
		return 0, errors.New("no id was generated")
	}
	var id int64
	err = result.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r DbSnippetRepository) LatestSnippets() (snp.Snippets, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	ORDER BY created DESC LIMIT 10`

	rows, err := r.Db.Query(stmt)

	if err != nil {
		return nil, err
	}

	// This should come after we check for an error
	// or the rows object could be nil
	defer rows.Close()

	snippets := snp.Snippets{}

	for rows.Next() {

		s := &snp.Snippet{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (r DbSnippetRepository) GetSnippet(id int64) (*snp.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE id = $1`

	row := r.Db.QueryRow(stmt, id)

	s := &snp.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
