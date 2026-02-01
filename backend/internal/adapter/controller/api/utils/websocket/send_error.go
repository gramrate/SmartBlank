package websocket

import "github.com/gorilla/websocket"

func (_ *WebSocket) SendError(ws *websocket.Conn, message string) {
	response := map[string]interface{}{
		"status":  "error",
		"message": message,
	}
	ws.WriteJSON(response)
}
