package websocket

import "github.com/gorilla/websocket"

func (_ *WebSocket) SendSuccess(ws *websocket.Conn, message string, data interface{}) {
	response := map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
	ws.WriteJSON(response)
}
