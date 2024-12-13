package utils

import "github.com/gorilla/websocket"

func SendResponse(conn *websocket.Conn, response interface{}) {
    resp := map[string]interface{}{
        "status": "success",
        "data":   response,
    }
    conn.WriteJSON(resp)
}

func SendError(conn *websocket.Conn, errorMessage string) {
    resp := map[string]interface{}{
        "status":  "error",
        "message": errorMessage,
    }
    conn.WriteJSON(resp)
}
