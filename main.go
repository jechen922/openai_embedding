package main

import (
	"fmt"
	"log"
	"net/http"
	"openai_golang/config"
	"openai_golang/openai/embedding"
	"openai_golang/src/database/mysql"
	"openai_golang/src/lib/file"
	"openai_golang/src/service"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	if err := config.EnvInit(); err != nil {
		log.Fatal("config.EnvInit()", err)
	}

	// 計算向量儲存至資料庫
	csvRecords := file.ReadCSVByFields("./resources/yile/yile.csv", "title", "heading", "content")
	sections := make([]embedding.Section, 0, len(csvRecords))
	for _, record := range csvRecords {
		sections = append(sections, embedding.Section{
			Title:   record["title"],
			Heading: record["heading"],
			Content: record["content"],
		})
	}
	conns := mysql.NewConnection()
	es := service.NewEmbedding(conns)
	if err := es.SaveSections(sections); err != nil {
		log.Fatal(fmt.Sprintf("embedding service save sections error: %v", err.Error()))
	}

	//websocket.Start()
}
