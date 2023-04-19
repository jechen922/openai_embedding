package contentrepo

import (
	"database/sql"
	"github.com/pgvector/pgvector-go"
	"openaigo/src/model/po"
)

type IContentRepo interface {
	AllByCategory(db *sql.DB, category string) ([]po.OpenAIContent, error)
}

func New() IContentRepo {
	return &contentRepo{}
}

type contentRepo struct{}

func (r *contentRepo) AllByCategory(db *sql.DB, category string) ([]po.OpenAIContent, error) {
	contents := make([]po.OpenAIContent, 0)
	rows, err := db.Query(`SELECT * FROM openai.contents WHERE category = $1;`, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		c := po.OpenAIContent{}
		var embedding pgvector.Vector
		if err := rows.Scan(
			&c.ID,
			&c.Category,
			&c.Heading,
			&c.Content,
			&c.Tokens,
			&embedding,
		); err != nil {
			return nil, err
		}
		c.Embedding = embedding.Slice()
		contents = append(contents, c)
	}
	return contents, nil
}
