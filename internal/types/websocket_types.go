package types

import "github.com/gorilla/websocket"

type SendResponseFunc func(*websocket.Conn, string, map[string]interface{})
type SendErrorFunc func(*websocket.Conn, string, map[string]interface{})
