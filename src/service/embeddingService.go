package service

import (
	"errors"
	"fmt"
	"openai_golang/openai/embedding"
	"openai_golang/src/database/mysql"
	"openai_golang/src/model/po"
	"openai_golang/src/repository"
	"strconv"
	"strings"
)

type IEmbeddingService interface {
	SaveSections(sections []embedding.Section) error
}

type embeddingService struct {
	mysqlDB mysql.IDB
}

func NewEmbedding(mysqlDB mysql.IDB) IEmbeddingService {
	return &embeddingService{mysqlDB: mysqlDB}
}

func (s *embeddingService) SaveSections(sections []embedding.Section) error {
	results, err := embedding.Create(sections...)
	if err != nil {
		return errors.New(fmt.Sprintf("create embedding error: %s", err.Error()))
	}

	embeddings := make([]po.Embedding, 0, len(results))
	for _, result := range results {
		vectors := make([]string, len(result.Vectors))
		for i, vector := range result.Vectors {
			vectors[i] = strconv.FormatFloat(vector, 'f', 10, 64)
		}
		embeddings = append(embeddings, po.Embedding{
			Title:   result.Title,
			Heading: result.Heading,
			Content: result.Content,
			Vectors: strings.Join(vectors, ";"),
		})
	}
	embeddingRepo := repository.EmbeddingRepo{}
	if err = embeddingRepo.Save(s.mysqlDB.Conn(mysql.DBOpenAI), embeddings...); err != nil {
		return errors.New(fmt.Sprintf("repo save error: %s", err.Error()))
	}
	return nil
}
