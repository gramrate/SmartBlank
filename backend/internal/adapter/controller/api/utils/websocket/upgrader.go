package websocket

import "github.com/gorilla/websocket"

func (ws *WebSocket) Upgrader() *websocket.Upgrader {
	return ws.wsUpgrader
}
