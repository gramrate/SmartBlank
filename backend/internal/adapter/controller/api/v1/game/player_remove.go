package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleRemovePlayer(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.RemovePlayerRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateRemovePlayerRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.RemovePlayer(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("remove player failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "player removed", game)
}
