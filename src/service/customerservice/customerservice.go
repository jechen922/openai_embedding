package customerservice

import (
	"database/sql"
	"errors"
	"fmt"
	"openaigo/openai/embedding"
	"openaigo/src/lib/calculator"
	"openaigo/src/model/po"
	"openaigo/src/repository"
	"openaigo/src/utils/openai"
	"sort"
)

type ICustomerService interface {
	Ask(question string) (string, error)
}

type customerService struct {
	postgresDB *sql.DB
	repo       repository.ICore
	ai         openai.IAI
}

func NewCustomer(postgresDB *sql.DB, repo repository.ICore, ai openai.IAI) ICustomerService {
	return &customerService{postgresDB: postgresDB, repo: repo, ai: ai}
}

func (cs *customerService) Ask(question string) (string, error) {
	intention, err := cs.ai.CustomerService().Classify(question)
	if err != nil {
		return "", err
	}
	switch intention {
	case "unknown":
		return "非常抱歉，我不太明白您的問題", nil
	}
	contents, err := cs.repo.Content().AllByCategory(cs.postgresDB, intention)
	if err != nil {
		return "", errors.New(fmt.Sprintf("repo AllByCategory error: %s", err.Error()))
	}

	qs, _ := embedding.Create(embedding.Section{
		Content: question,
	})
	// 計算向量之間的相似度
	vectorEmbeddingMap := make(map[float32]po.OpenAIContent, len(contents))
	similarities := make([]float32, 0, len(contents))
	for _, c := range contents {
		similarity := calculator.CosineSimilarity(qs.Vectors, c.Embedding)
		//fmt.Printf("\"%s\" 和 \"%s\" 之間的相似度: %f\n", qs[0].Content, e.Heading, similarity)
		vectorEmbeddingMap[similarity] = c
		similarities = append(similarities, similarity)
	}
	//// 使用 sort 函數對 slice 進行排序
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i] > similarities[j]
	})

	// 取得前三個最大值
	provideAnswers := make([]string, 0)
	for i := 0; i < 3; i++ {
		if len(similarities) >= i+1 {
			fmt.Printf("title: %v, heading: %v, simi: %v\n",
				vectorEmbeddingMap[similarities[i]].Category,
				vectorEmbeddingMap[similarities[i]].Heading,
				similarities[i])
			provideAnswers = append(provideAnswers, vectorEmbeddingMap[similarities[i]].Content)
		}
	}
	ans := cs.ai.CustomerService().Answer(question, provideAnswers...)
	return ans, nil
}
