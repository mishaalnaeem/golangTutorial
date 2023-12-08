package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

var clients = make(map[string]*websocket.Conn)

func main() {
	http.HandleFunc("/echo", handlerConnection)
	http.HandleFunc("/", serverStaticFile)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err!=nil {
		log.Fatal(err)
	}

	clients[conn.RemoteAddr().String()] = conn

	go handlerClient(conn)
}

func deleteConection(conn *websocket.Conn) {
	delete(clients, conn.RemoteAddr().String())
}

func handlerClient(conn *websocket.Conn) {

	//for closing connection when functions exits
	defer func() {
		deleteConection(conn)
		conn.Close()
	}()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err!=nil {
			log.Fatal(err)
		}

		fmt.Printf("%s: %s",conn.RemoteAddr().String(), string(msg))
		go broadcaster(msgType, msg)
	}
}

func broadcaster(msgType int, msg []byte) {
	for _, client := range clients {
		if err := client.WriteMessage(msgType, msg); err!=nil {
			log.Fatal(err)
		}
	}
}

func serverStaticFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}