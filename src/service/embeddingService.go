package service

import (
	"openaigo/openai/chat"
	"openaigo/src/database/mysql"
)

type IEmbeddingService interface {
	AnswerByChat(question string) (string, error)
	//SaveSections(sections []embedding.Section) error
}

type embeddingService struct {
	mysqlDB mysql.IDB
}

func NewEmbedding(mysqlDB mysql.IDB) IEmbeddingService {
	return &embeddingService{mysqlDB: mysqlDB}
}

func (s *embeddingService) AnswerByChat(question string) (string, error) {
	title := chat.GetTitle(question)
	switch title {
	//case "個人資料":
	//	return chat.AnswerPersonalQuestion(question), nil
	case "unknown":
		return "非常抱歉，我不太明白您的問題", nil
	}
	//embeddingRepo := repository.EmbeddingRepo{}
	//embeddings, err := embeddingRepo.AllByTitle(s.mysqlDB.Conn(mysql.DBOpenAI), title)
	//if err != nil {
	//	return "", errors.New(fmt.Sprintf("repo All error: %s", err.Error()))
	//}
	//
	//qs, _ := embedding.Create(embedding.Section{
	//	Content: question,
	//})

	// 計算向量之間的相似度
	//vectorEmbeddingMap := make(map[float32]po.Embedding, len(embeddings))
	//similarities := make([]float32, 0, len(embeddings))
	//for _, e := range embeddings {
	//	vectorList := strings.Split(e.Vectors, ";")
	//	vectors := make([]float32, len(vectorList))
	//	for i, v := range vectorList {
	//		vectors[i], _ = strconv.ParseFloat(v, 32)
	//	}
	//
	//	similarity := calculator.CosineSimilarity(qs[0].Vectors, vectors)
	//	//fmt.Printf("\"%s\" 和 \"%s\" 之間的相似度: %f\n", qs[0].Content, e.Heading, similarity)
	//	vectorEmbeddingMap[similarity] = e
	//	similarities = append(similarities, similarity)
	//}
	//// 使用 sort 函數對 slice 進行排序
	//sort.Slice(similarities, func(i, j int) bool {
	//	return similarities[i] > similarities[j]
	//})

	// 取得前三個最大值
	answers := make([]string, 0)
	//for i := 0; i < 3; i++ {
	//	if len(similarities) >= i+1 {
	//		fmt.Printf("title: %v, heading: %v, simi: %v\n",
	//			vectorEmbeddingMap[similarities[i]].Title,
	//			vectorEmbeddingMap[similarities[i]].Heading,
	//			similarities[i])
	//		answers = append(answers, vectorEmbeddingMap[similarities[i]].Content)
	//	}
	//}
	ans := chat.Chat(question, answers...)
	return ans, nil
}

//func (s *embeddingService) SaveSections(sections []embedding.Section) error {
//	results, err := embedding.Create(sections...)
//	if err != nil {
//		return errors.New(fmt.Sprintf("create embedding error: %s", err.Error()))
//	}
//
//	embeddings := make([]po.Embedding, 0, len(results))
//	for _, result := range results {
//		vectors := make([]string, len(result.Vectors))
//		for i, vector := range result.Vectors {
//			vectors[i] = strconv.FormatFloat(vector, 'f', 10, 64)
//		}
//		embeddings = append(embeddings, po.Embedding{
//			Title:   result.Title,
//			Heading: result.Heading,
//			Content: result.Content,
//			Vectors: strings.Join(vectors, ";"),
//		})
//	}
//	embeddingRepo := repository.EmbeddingRepo{}
//	if err = embeddingRepo.Save(s.mysqlDB.Conn(mysql.DBOpenAI), embeddings...); err != nil {
//		return errors.New(fmt.Sprintf("repo save error: %s", err.Error()))
//	}
//	return nil
//}
