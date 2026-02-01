package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleStartVote(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.StartVoteRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateStartVoteRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.StartVote(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("start vote failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "vote started", game)
}
