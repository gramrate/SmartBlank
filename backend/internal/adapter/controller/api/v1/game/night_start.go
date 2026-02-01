package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleStartNight(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.StartNightRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateStartNightRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.StartNight(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("start night failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "night started", game)
}
