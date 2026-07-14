package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

// Snippet struct
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgx.Conn
}

func (s *SnippetModel) Insert(title, content string, expires int) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires)
            VALUES ($1, $2, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc' + ($3 * INTERVAL '1 day'))
            RETURNING id`

	var id int
	err := s.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SnippetModel) Get(id int) (Snippet, error) {

	return Snippet{}, nil
}

func (s *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
