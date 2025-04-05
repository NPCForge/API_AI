package handlers

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Server struct {
	Conns      map[*websocket.Conn]bool
	Mutex      sync.Mutex
	IsBlocking bool
	BlockMutex sync.RWMutex
}

var WS = &Server{
	Conns: make(map[*websocket.Conn]bool),
}
