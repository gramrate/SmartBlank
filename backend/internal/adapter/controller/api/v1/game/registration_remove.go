package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleRemoveRegistrationPlayer(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.RemoveRegistrationPlayerRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateRemoveRegistrationPlayer(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.RemoveRegistrationPlayer(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("remove registration player failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "registration player removed", game)
}
