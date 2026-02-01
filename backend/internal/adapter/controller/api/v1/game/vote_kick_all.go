package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleKickAllVote(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.KickAllVoteRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateKickAllVoteRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.KickAllVote(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("kick all vote failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "kick all vote updated", game)
}
