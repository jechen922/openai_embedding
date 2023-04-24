package contentrepo

import (
	"database/sql"
	"github.com/pgvector/pgvector-go"
	"openaigo/src/model/po"
)

type IContentRepo interface {
	All(db *sql.DB) ([]po.OpenAIContent, error)
	AllByCategory(db *sql.DB, category string) ([]po.OpenAIContent, error)
	GetByHeading(db *sql.DB, heading string) (po.OpenAIContent, error)
	AllCategories(db *sql.DB) ([]string, error)
	AllHeadings(db *sql.DB) ([]string, error)
}

func New() IContentRepo {
	return &contentRepo{}
}

type contentRepo struct{}

func (r *contentRepo) All(db *sql.DB) ([]po.OpenAIContent, error) {
	contents := make([]po.OpenAIContent, 0)
	rows, err := db.Query(`SELECT * FROM openai.contents;`)
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

func (r *contentRepo) GetByHeading(db *sql.DB, heading string) (po.OpenAIContent, error) {
	c := po.OpenAIContent{}
	row := db.QueryRow(`SELECT * FROM openai.contents WHERE heading = $1;`, heading)

	var embedding pgvector.Vector
	if err := row.Scan(
		&c.ID,
		&c.Category,
		&c.Heading,
		&c.Content,
		&c.Tokens,
		&embedding,
	); err != nil {
		return po.OpenAIContent{}, err
	}
	c.Embedding = embedding.Slice()

	return c, nil
}

func (r *contentRepo) AllCategories(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`SELECT category FROM openai.contents GROUP BY category;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := make([]string, 0)
	for rows.Next() {
		var category string
		if err := rows.Scan(
			&category,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *contentRepo) AllHeadings(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`SELECT heading FROM openai.contents GROUP BY heading;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	headings := make([]string, 0)
	for rows.Next() {
		var heading string
		if err := rows.Scan(
			&heading,
		); err != nil {
			return nil, err
		}
		headings = append(headings, heading)
	}
	return headings, nil
}
