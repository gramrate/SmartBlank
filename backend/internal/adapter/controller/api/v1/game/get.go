package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleGet(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.LobbyIDRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateLobbyIDRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.Get(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("get game failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "game state", game)
}
