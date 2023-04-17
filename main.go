package main

import (
	"log"
	"openaigo/config"
	"openaigo/src/tools/websocket"
	"openaigo/src/tools/wire"
)

func main() {
	app, err := wire.InitializeApplication(config.New())
	if err != nil {
		log.Fatal(err)
	}
	go websocket.Start()
	fiberRoute := app.Fiber.Static("/", "./view/")
	app.Router.Set(fiberRoute)
	err = app.Fiber.Listen(":8081")
	if err != nil {
		panic(err) // 若啟動伺服器失敗，拋出錯誤訊息
	}
}
