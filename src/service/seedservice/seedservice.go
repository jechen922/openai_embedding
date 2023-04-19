package seedservice

import (
	"database/sql"
	"errors"
	"fmt"
	"openaigo/openai/embedding"
	"openaigo/src/model/po"
	"openaigo/src/repository"
)

type ISeedService interface {
	SaveSections(sections []embedding.Section) error
}

type seedService struct {
	postgresDB *sql.DB
	repo       repository.ICore
}

func NewSeed(postgresDB *sql.DB, repo repository.ICore) ISeedService {
	return &seedService{postgresDB: postgresDB, repo: repo}
}

func (s *seedService) SaveSections(sections []embedding.Section) error {
	for _, section := range sections {
		result, err := embedding.Create(section)
		if err != nil {
			return errors.New(fmt.Sprintf("create embedding error: %s", err.Error()))
		}
		openaiContent := po.OpenAIContent{
			Category:  result.Title,
			Heading:   result.Heading,
			Content:   result.Content,
			Tokens:    result.Tokens,
			Embedding: result.Vectors,
		}
		if err = s.repo.Seed().Save(s.postgresDB, openaiContent); err != nil {
			return errors.New(fmt.Sprintf("repo save error: %s", err.Error()))
		}
	}
	return nil
}
