package ws

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

type ChallengeReq struct {
	ID         string                 `json:"i"`
	Command    string                 `json:"c"`
	Parameters map[string]interface{} `json:"p"`
}

type ChallengeRes struct {
	ID      string                 `json:"i"`
	Status  int16                  `json:"s"`
	Message map[string]interface{} `json:"m,omitempty"`
	Error   string                 `json:"e,omitempty"`
}

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	conn, deferFunc, err := upgradeConnection(w, r)
	if err != nil {
		log.Print("failed to upgrade websocket:", err)
		return
	}
	defer deferFunc()

	connectionID := uuid.NewV4().String()
	log.Printf("client connected. connectionID: %s", connectionID)

	for {
		request := ChallengeReq{}
		err := conn.ReadJSON(&request)
		if err != nil {
			// handle client disconnection
			if _, ok := err.(*websocket.CloseError); ok {
				log.Printf("disconnected from client. connectionID: %s", connectionID)
				break
			}

			// send error message to client
			_ = conn.WriteJSON(ChallengeRes{
				ID:     "",
				Status: -1,
				Error:  err.Error(),
			})
			continue
		}

		log.Printf("[RECV] %v", request)

		if err := conn.WriteJSON(ChallengeRes{
			ID:      "",
			Status:  0,
			Message: map[string]interface{}{"foo": "bar"},
		}); err != nil {
			log.Print("failed to write:", err)
			break
		}
	}
}
