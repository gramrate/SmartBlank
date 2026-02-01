package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleAddRegistrationPlayer(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.AddRegistrationPlayerRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateAddRegistrationPlayer(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.AddRegistrationPlayer(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("add registration player failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "registration player added", game)
}
