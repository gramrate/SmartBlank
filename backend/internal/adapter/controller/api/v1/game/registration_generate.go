package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleGenerateSeating(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.GenerateSeatingRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateGenerateSeatingRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.GenerateSeating(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("generate seating failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "seating generated", game)
}
