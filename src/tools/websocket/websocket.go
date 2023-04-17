package websocket

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func Start() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	addr := flag.String("addr", ":8080", "http service address")
	fmt.Println("websocket server starting on ", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
