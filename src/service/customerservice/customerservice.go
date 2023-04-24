package customerservice

import (
	"database/sql"
	"fmt"
	"math"
	"openaigo/openai/embedding"
	"openaigo/src/lib/calculator"
	"openaigo/src/model/po"
	"openaigo/src/repository"
	"openaigo/src/utils/openai"
	"sort"

	"github.com/agnivade/levenshtein"
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

func isSemanticSimilar(str1, str2 string) bool {
	distance := levenshtein.ComputeDistance(str1, str2)
	similarity := 1.0 - float64(distance)/float64(math.Max(float64(len(str1)), float64(len(str2))))
	// 可根据实际情况设定相似性阈值，例如 0.8
	threshold := 0.7
	//if similarity >= threshold {
	fmt.Println("similarity", similarity, str1, str2)
	//}
	return similarity >= threshold
}

func (cs *customerService) Ask(question string) (string, error) {
	fmt.Println("ask question:", question)
	//intention, err := cs.ai.CustomerService().Classify(question)
	//if err != nil {
	//	return "", err
	//}
	//switch intention {
	//case "unknown":
	//	return "非常抱歉，我不太明白您的問題", nil
	//}

	contents := []po.OpenAIContent{}
	categories, _ := cs.repo.Content().AllCategories(cs.postgresDB)
	for _, category := range categories {
		if isSemanticSimilar(question, category) {
			ctns, _ := cs.repo.Content().AllByCategory(cs.postgresDB, category)
			contents = append(contents, ctns...)
		}
	}
	if len(contents) == 0 {
		headings, _ := cs.repo.Content().AllHeadings(cs.postgresDB)
		for _, heading := range headings {
			if isSemanticSimilar(question, heading) {
				ctn, _ := cs.repo.Content().GetByHeading(cs.postgresDB, heading)
				contents = append(contents, ctn)
			}
		}
	}
	if len(contents) == 0 {
		fmt.Println("no contents")
		return "no contents", nil
	}
	//contents, err := cs.repo.Content().All(cs.postgresDB)
	//if err != nil {
	//	return "", errors.New(fmt.Sprintf("repo AllByCategory error: %s", err.Error()))
	//}
	//fmt.Println(contents)
	qs, _ := embedding.Content(embedding.Section{
		Content: question,
	})
	// 計算向量之間的相似度
	idContentMap := make(map[int64]po.OpenAIContent, len(contents))
	vectorEmbeddingMap := make(map[int64]float32, len(contents))
	similarities := make([]float32, 0, len(contents))
	for _, c := range contents {
		idContentMap[c.ID] = c
		//fmt.Println("contents:", c.Category, c.Heading)
		similarity := calculator.CosineSimilarity(qs.Vectors, c.Embedding)
		//fmt.Printf("\"%s\" 和 \"%s\" 之間的相似度: %f\n", qs[0].Content, e.Heading, similarity)
		vectorEmbeddingMap[c.ID] = similarity
		similarities = append(similarities, similarity)
		//fmt.Println("similarity", similarity, c.Category, c.Heading)
	}
	//// 使用 sort 函數對 slice 進行排序
	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i] > similarities[j]
	})
	//for _, similarity := range similarities {
	//	fmt.Printf("title: %v, heading: %v, simi: %v, token:%d\n",
	//		vectorEmbeddingMap[similarity].Category,
	//		vectorEmbeddingMap[similarity].Heading,
	//		similarity,
	//		vectorEmbeddingMap[similarity].Tokens,
	//	)
	//}

	// 取得前三個最大值
	tokens := 0
	provideAnswers := make([]string, 0)
	provideIDMap := make(map[int64]struct{})
	for i := range similarities {
		for id, simi := range vectorEmbeddingMap {
			if similarities[i] == simi {
				tokens += idContentMap[id].Tokens
				if tokens > 2000 {
					break
				}
				if _, exist := provideIDMap[id]; exist {
					continue
				}
				//fmt.Printf("!!!gotcha title: %v, heading: %v, simi: %v, token:%d, content: %v\n",
				//	idContentMap[id].Category,
				//	idContentMap[id].Heading,
				//	similarities[i],
				//	idContentMap[id].Tokens,
				//	idContentMap[id].Content,
				//)
				fmt.Printf("!!!gotcha, ID:%+v title: %v, heading: %v, simi: %v, token:%d\n",
					idContentMap[id].ID,
					idContentMap[id].Category,
					idContentMap[id].Heading,
					similarities[i],
					idContentMap[id].Tokens,
				)
				provideAnswers = append(provideAnswers, idContentMap[id].Content)
				provideIDMap[id] = struct{}{}
			}
		}
	}
	ans := cs.ai.CustomerService().Answer(question, provideAnswers...)
	return ans, nil
}
