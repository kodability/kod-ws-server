package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

// StartServer starts a websocket server.
func StartServer(port uint16) {
	http.HandleFunc("/", echoHandler)

	addr := fmt.Sprintf("localhost:%d", port)
	log.Printf("starting server on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}