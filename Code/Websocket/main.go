package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func reader(conn *websocket.Conn) {
	messageType, p, err := conn.ReadMessage()
	for {
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//Avoid CORS error
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client succesfully connected...")

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("vim-go")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
