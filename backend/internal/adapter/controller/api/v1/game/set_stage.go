package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleSetStage(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.SetStageRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateSetStageRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.SetStage(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("set stage failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "stage updated", game)
}
