package seedservice

import (
	"database/sql"
	"errors"
	"fmt"
	"openaigo/openai/embedding"
	"openaigo/src/model/po"
	"openaigo/src/repository/seedrepo"
)

type ISeedService interface {
	SaveSections(sections []embedding.Section) error
}

type seedService struct {
	postgresDB *sql.DB
}

func NewSeed(postgresDB *sql.DB) ISeedService {
	return &seedService{postgresDB: postgresDB}
}

func (s *seedService) SaveSections(sections []embedding.Section) error {
	// 計算向量儲存至資料庫
	//csvRecords := file.ReadCSVByFields("../../../resources/yile/yile.csv", "title", "heading", "content")
	//sections := make([]embedding.Section, 0, len(csvRecords))
	//for _, record := range csvRecords {
	//	sections = append(sections, embedding.Section{
	//		Title:   record["title"],
	//		Heading: record["heading"],
	//		Content: record["content"],
	//	})
	//}

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
		seedRepo := seedrepo.New()
		if err = seedRepo.Save(s.postgresDB, openaiContent); err != nil {
			return errors.New(fmt.Sprintf("repo save error: %s", err.Error()))
		}
	}
	return nil
}
