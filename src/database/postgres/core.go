package postgres

//
//import (
//	"os"
//	"sort"
//
//	"github.com/pgvector/pgvector-go"
//	"github.com/wangyuche/aics/src/gpt"
//	"github.com/wangyuche/aics/src/utils"
//	"github.com/wangyuche/goutils/log"
//	"github.com/wangyuche/goutils/sqllib"
//)
//
//var pg sqllib.ISQL
//
//func init() {
//	pg = sqllib.New(sqllib.SQLType(os.Getenv("SQLType")))
//	pg.WR().Init(os.Getenv("SQLURL"), 600, 1800, 2, 10)
//	pg.R().Init(os.Getenv("SQLURL"), 600, 1800, 2, 10)
//}
//
//func CSVtoTrain(file string) error {
//	csvlist, err := utils.GetCSVList(file)
//	if err != nil {
//		return err
//	}
//	for _, line := range csvlist {
//		log.Debug(line[0])
//		l, err := utils.GetTokenLength(line[0])
//		if err != nil {
//			return err
//		}
//		vectors, err := gpt.EmbeddingTrain([]string{line[0]})
//		if err != nil {
//			return err
//		}
//		ContenttoDB(line[0], l, vectors)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func ContenttoDB(content string, token int, vectors []float32) error {
//	db, err := pg.WR().GetTX()
//	if err != nil {
//		log.Fail(err.Error())
//		return err
//	}
//
//	var id int
//	err = db.QueryRow("INSERT INTO openai.contents (content, tokens) VALUES ($1, $2) returning id;", content, token).Scan(&id)
//	if err != nil {
//		log.Fail(err.Error())
//		return err
//	}
//	_, err = db.Exec("INSERT INTO openai.vectors (idcontents, embedding) VALUES ($1, $2)", id, pgvector.NewVector(vectors))
//	if err != nil {
//		db.Rollback()
//		log.Fail(err.Error())
//		return err
//	}
//	db.Commit()
//	return nil
//}
//
//type Distances struct {
//	Dis     float32
//	Content string
//}
//
//func AnswerQuestion(question string) (string, error) {
//	var dis []Distances = make([]Distances, 0)
//	vectors, err := gpt.EmbeddingTrain([]string{question})
//	if err != nil {
//		return "", err
//	}
//	db := pg.R().GetDB()
//	rows, err := db.Query("SELECT a.id, a.content, b.embedding FROM openai.contents as a INNER JOIN openai.vectors as b ON a.id = b.idcontents")
//	if err != nil {
//		log.Fail(err.Error())
//		return "", err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var Id int
//		var content string
//		var embedding pgvector.Vector
//		var d Distances
//		rows.Scan(&Id, &content, &embedding)
//		d.Dis = utils.Cosine(vectors, embedding.Slice())
//		d.Content = content
//		dis = append(dis, d)
//	}
//	sort.SliceStable(dis, func(i, j int) bool {
//		return dis[i].Dis < dis[j].Dis
//	})
//	ans, err := gpt.CreateCompletion(`Answer the question based on the context below, and if the question can't be answered based on the context, say "I don't know",
//	                                  Context: ` + dis[len(dis)-1].Content + `
//									  Question: ` + question + `
//									  Answer:`)
//	if err != nil {
//		return "", err
//	}
//	return ans, nil
//}
