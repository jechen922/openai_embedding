package postgres

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("SQLURL"))
	if err != nil {
		log.Fatal("sql.Open:", err)
	}
	//defer db.Close()

	// 測試連線
	err = db.Ping()
	if err != nil {
		log.Fatal("db.Ping:", err)
	}
	return db, nil

	//pg := sqllib.New(sqllib.SQLType(os.Getenv("SQLType")))
	//pg.WR().Init(os.Getenv("SQLURL"), 600, 1800, 2, 10)
	//pg.R().Init(os.Getenv("SQLURL"), 600, 1800, 2, 10)
	//
	//db, err := pg.WR().GetTX()
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return err
	//}
	//
	//var id int
	//err = db.QueryRow("INSERT INTO openai.contents (content, tokens) VALUES ($1, $2) returning id;", content, token).Scan(&id)
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return err
	//}
	//_, err = db.Exec("INSERT INTO openai.vectors (idcontents, embedding) VALUES ($1, $2)", id, pgvector.NewVector(vectors))
	//if err != nil {
	//	db.Rollback()
	//	log.Fatal(err.Error())
	//	return err
	//}
	//db.Commit()
	//return nil
}
