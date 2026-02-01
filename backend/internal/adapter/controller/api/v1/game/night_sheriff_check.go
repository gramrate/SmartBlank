package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleSheriffCheck(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.SheriffCheckRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateSheriffCheckRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.SheriffCheck(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("sheriff check failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "sheriff check done", game)
}
