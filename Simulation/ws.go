package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader    = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients     = make(map[*websocket.Conn]bool)
	broadcastCh = make(chan []byte)
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		debug.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// Register the new client
	clients[conn] = true
	log.Println("New client connected")
	debug.Println("New client connected")

	// Handle client messages (optional, just to keep the connection open)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			debug.Println("Read error:", err)
			delete(clients, conn)
			break
		}
	}
}

func periodicBroadcast() {
	for {
		select {
		case <-Ticker.C:
			for _, v := range Veheicles {
				mapD := map[string]int{"id": v.id, "x": v.position[0], "y": v.position[1]}
				mapB, _ := json.Marshal(mapD)

				// Send updates to all connected clients
				for client := range clients {
					broadcastLog.Println("to ", client.LocalAddr().String(), " :", string(mapB))
					err := client.WriteMessage(websocket.TextMessage, mapB)
					if err != nil {
						debug.Println("Write error:", err)
						broadcastLog.Println("Write error:", err)
						client.Close()
						delete(clients, client) // Remove the client if there is an error
					}
				}
			}
		}
	}
}
