package models

import (
	"github.com/gorilla/websocket"
)

type WebSocketDispatcher struct {
	Name      string
	Handler   func(*websocket.Conn, []byte, func(*websocket.Conn, interface{}), func(*websocket.Conn, string))
	Protected bool
}
