package postgresql

import (
	"database/sql"
	"errors"

	"github.com/Sleenjep/snippetbox-proj/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `
		INSERT INTO snippets (title, content, created, expires)
        VALUES ($1, $2, NOW(), NOW() + ($3 || 'days')::interval)
        RETURNING id;
	`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}

	query := `
		SELECT id, title, content, created, expires
		FROM snippets
		WHERE expires > NOW() AT TIME ZONE 'UTC' AND id = $1
	`

	err := m.DB.QueryRow(query, id).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&s.Created,
		&s.Expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
