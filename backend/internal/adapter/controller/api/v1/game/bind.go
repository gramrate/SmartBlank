package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleBind(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.LobbyIDRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateLobbyIDRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	if _, err := h.gameService.Get(context.Background(), &req); err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("bind failed: %w", err).Error())
		return
	}
	h.bindLobbyID(ws, req.LobbyID)
	h.wsUtils.SendSuccess(ws, "lobby bound", map[string]string{"lobby_id": req.LobbyID})
}
