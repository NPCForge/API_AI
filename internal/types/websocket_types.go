package types

import "github.com/gorilla/websocket"

type SendResponseFunc func(*websocket.Conn, string, string, map[string]interface{})
type SendErrorFunc func(*websocket.Conn, string, string, map[string]interface{})
