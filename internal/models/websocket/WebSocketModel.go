package websocketModels

import (
	"github.com/gorilla/websocket"
)

type WebSocketMessage struct {
	Action string `json:"action"`
}

type WebSocketDispatcher struct {
	Name      string
	Handler   func(*websocket.Conn, []byte, func(*websocket.Conn, interface{}), func(*websocket.Conn, string))
	Protected bool
}
