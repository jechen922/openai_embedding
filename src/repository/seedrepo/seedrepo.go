package seedrepo

import (
	"database/sql"
	"openaigo/src/model/po"

	"github.com/pgvector/pgvector-go"
)

type ISeedRepo interface {
	SaveCategory(db *sql.DB, c po.OpenAICategory) error
	SaveContent(db *sql.DB, c po.OpenAIContent) error
}

func New() ISeedRepo {
	return &seedRepo{}
}

type seedRepo struct{}

//`
//		INSERT INTO openai.contents (
//			id, category, heading, content, tokens, embedding
//		) VALUES (
//	  		$1, $2, $3, $4, $5, $6
//		) ON CONFLICT (category, heading, content) DO UPDATE
//		SET content = EXCLUDED.content,
//		    tokens = EXCLUDED.tokens,
//		    embedding = EXCLUDED.embedding;`,

func (r *seedRepo) SaveCategory(db *sql.DB, c po.OpenAICategory) error {
	_, err := db.Exec(
		`
		INSERT INTO openai.category (
			 category, tokens, embeddings
		) VALUES (
	  		 $1, $2, $3
		);`,
		c.Category,
		c.Tokens,
		pgvector.NewVector(c.Embedding),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *seedRepo) SaveContent(db *sql.DB, c po.OpenAIContent) error {
	_, err := db.Exec(
		`
		INSERT INTO openai.contents (
			 category, heading, content, tokens, embedding
		) VALUES (
	  		 $1, $2, $3, $4, $5
		);`,
		c.Category,
		c.Heading,
		c.Content,
		c.Tokens,
		pgvector.NewVector(c.Embedding),
	)
	if err != nil {
		return err
	}
	return nil
}
