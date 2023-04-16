package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"openaigo/config"
	"openaigo/src/router"
	"openaigo/websocket"
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
	//conns := mysql.NewConnection()
	//es := service.NewEmbedding(conns)

	// 計算向量儲存至資料庫
	//csvRecords := file.ReadCSVByFields("./resources/yile/yile.csv", "title", "heading", "content")
	//sections := make([]embedding.Section, 0, len(csvRecords))
	//for _, record := range csvRecords {
	//	sections = append(sections, embedding.Section{
	//		Title:   record["title"],
	//		Heading: record["heading"],
	//		Content: record["content"],
	//	})
	//}
	//if err := es.SaveSections(sections); err != nil {
	//	log.Fatal(fmt.Sprintf("embedding service save sections error: %v", err.Error()))
	//}

	//question := "你叫什麼名字?在哪裡上班？身高體重多少？擔任什麼職位？你們公司是做什麼的？"
	//question := "可以介紹一下你們公司嗎"
	//question := "如何申請出差補助？"
	//ans, _ := es.AnswerByChat(question)
	//fmt.Println("question:", question)
	//fmt.Println("answer:", ans)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
	})

	app.Static("/", "./view/")
	go websocket.Start()

	router.Set(app)
	err := app.Listen(":8081")
	if err != nil {
		panic(err) // 若啟動伺服器失敗，拋出錯誤訊息
	}
}
