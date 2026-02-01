package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleUpdateMusic(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.UpdateMusicRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateUpdateMusicRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.UpdateMusic(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("update music failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "music updated", game)
}
