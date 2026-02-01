package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleSetBestTurn(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.SetBestTurnRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateSetBestTurnRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.SetBestTurn(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("set best turn failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "best turn set", game)
}
