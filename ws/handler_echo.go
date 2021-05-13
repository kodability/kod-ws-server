package ws

import (
	"log"
	"net/http"
)

func echoHandler(writer http.ResponseWriter, req *http.Request) {
	conn, deferFunc, err := upgradeConnection(writer, req)
	if err != nil {
		log.Print("failed to upgrade websocket:", err)
		return
	}
	defer deferFunc()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Print("failed to read:", err)
			break
		}

		log.Printf("[RECV] %s", msg)
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Print("failed to write:", err)
			break
		}
	}
}
