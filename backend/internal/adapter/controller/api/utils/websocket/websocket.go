package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	wsUpgrader *websocket.Upgrader
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		wsUpgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Настройте properly для продакшена
			},
		},
	}
}
