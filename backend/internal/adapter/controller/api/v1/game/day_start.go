package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleStartDay(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.StartDayRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateStartDayRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.StartDay(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("start day failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "day started", game)
}
