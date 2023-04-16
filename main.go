package main

import (
	"log"
	"net/http"
	"openaigo/config"
	"openaigo/src/tools/wire"
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
	app, err := wire.InitializeApplication(config.New())
	if err != nil {
		log.Fatal(err)
	}
	app.Fiber.Static("/", "./view/")
	go websocket.Start()
	err = app.Fiber.Listen(":8081")
	if err != nil {
		panic(err) // 若啟動伺服器失敗，拋出錯誤訊息
	}

	//question := "你叫什麼名字?在哪裡上班？身高體重多少？擔任什麼職位？你們公司是做什麼的？"
	//question := "可以介紹一下你們公司嗎"
	//question := "如何申請出差補助？"
	//ans, _ := es.AnswerByChat(question)
	//fmt.Println("question:", question)
	//fmt.Println("answer:", ans)
}
