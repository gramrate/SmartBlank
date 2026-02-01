package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleApplyNightResults(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.ApplyNightResultsRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateApplyNightResultsRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.ApplyNightResults(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("apply night results failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "night results applied", game)
}
