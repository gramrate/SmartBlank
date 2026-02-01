package service_provider

import (
	wsUtils "backend/internal/adapter/controller/api/utils/websocket"
)

func (s *ServiceProvider) WebSocketUtils() *wsUtils.WebSocket {
	if s.wsUtils == nil {
		s.wsUtils = wsUtils.NewWebSocket()
	}

	return s.wsUtils
}
