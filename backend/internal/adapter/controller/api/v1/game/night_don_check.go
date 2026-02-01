package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleDonCheck(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.DonCheckRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateDonCheckRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.DonCheck(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("don check failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "don check done", game)
}
