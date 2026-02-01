package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleSetCard(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.SetCardRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateSetCardRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.SetCard(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("set card failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "card updated", game)
}
