package models

import (
	"context"
	"errors"
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

func (m *SnippetModel) Get(id int) (Snippet, error) {
	// 1. Меняем синтаксис SQL:
	// - UTC_TIMESTAMP() заменяем на NOW() AT TIME ZONE 'utc'
	// - Плейсхолдер ? заменяем на $1
	stmt := `SELECT id, title, content, created, expires FROM snippets
             WHERE expires > NOW() AT TIME ZONE 'utc' AND id = $1`

	// Инициализируем пустую структуру Snippet
	var s Snippet

	// 2. Вызываем QueryRow на пуле pgx, обязательно передавая контекст.
	// Метод Scan() вызываем цепочкой сразу после QueryRow — это стандартный и лаконичный паттерн.
	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&s.Created,
		&s.Expires,
	)

	if err != nil {
		// 3. Вместо sql.ErrNoRows проверяем ошибку pgx.ErrNoRows
		if errors.Is(err, pgx.ErrNoRows) {
			return Snippet{}, ErrSnippetNotFound
		}
		return Snippet{}, err
	}

	// Если всё прошло успешно, возвращаем заполненную структуру
	return s, nil
}

func (s *SnippetModel) Latest() ([]Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets where expires > NOW() ORDER BY created DESC LIMIT 10`
	rows, err := s.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []Snippet
	for rows.Next() {
		var s Snippet
		err = rows.Scan(
			&s.ID,
			&s.Title,
			&s.Content,
			&s.Created,
			&s.Expires)
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
