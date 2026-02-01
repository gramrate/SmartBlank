package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleResolveVote(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.ResolveVoteRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateResolveVoteRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.ResolveVote(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("resolve vote failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "vote resolved", game)
}
