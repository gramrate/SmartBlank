package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleSwapRegistrationPositions(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.SwapRegistrationPositionsRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateSwapRegistrationPositions(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.SwapRegistrationPositions(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("swap registration positions failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "registration positions swapped", game)
}
