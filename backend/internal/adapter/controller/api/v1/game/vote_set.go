package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleSetVote(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.SetVoteRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateSetVoteRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.SetVote(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("set vote failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "vote updated", game)
}
