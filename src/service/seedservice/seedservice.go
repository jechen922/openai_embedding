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
		//categoryResult, err := embedding.Category(po.OpenAICategory{
		//	Category: section.Title + "-" + section.Heading,
		//})
		//if err != nil {
		//	return errors.New(fmt.Sprintf("create embedding error: %s", err.Error()))
		//}
		//openaiCategory := po.OpenAICategory{
		//	Category:  section.Title + "-" + section.Heading,
		//	Tokens:    categoryResult.Tokens,
		//	Embedding: categoryResult.Embedding,
		//}
		//if err = s.repo.Seed().SaveCategory(s.postgresDB, openaiCategory); err != nil {
		//	return errors.New(fmt.Sprintf("repo SaveCategory error: %s", err.Error()))
		//}

		contentResult, err := embedding.Content(section)
		if err != nil {
			return errors.New(fmt.Sprintf("create embedding error: %s", err.Error()))
		}
		openaiContent := po.OpenAIContent{
			Category:  contentResult.Title,
			Heading:   contentResult.Heading,
			Content:   contentResult.Content,
			Tokens:    contentResult.Tokens,
			Embedding: contentResult.Vectors,
		}
		//fmt.Println(openaiContent.Category, openaiContent.Heading, openaiContent.Content)
		if err = s.repo.Seed().SaveContent(s.postgresDB, openaiContent); err != nil {
			return errors.New(fmt.Sprintf("repo SaveContent error: %s", err.Error()))
		}
	}
	return nil
}
