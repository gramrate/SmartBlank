package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleMafiaTarget(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.MafiaTargetRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateMafiaTargetRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.MafiaTarget(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("mafia target failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "mafia target selected", game)
}
