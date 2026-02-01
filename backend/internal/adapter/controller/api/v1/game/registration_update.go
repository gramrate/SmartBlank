package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleUpdateRegistrationPlayer(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.UpdateRegistrationPlayerRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateUpdateRegistrationPlayer(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.UpdateRegistrationPlayer(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("update registration player failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "registration player updated", game)
}
