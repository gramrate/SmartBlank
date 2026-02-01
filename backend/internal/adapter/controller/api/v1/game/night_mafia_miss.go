package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleMafiaMiss(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.MafiaMissRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateMafiaMissRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.MafiaMiss(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("mafia miss failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "mafia miss", game)
}
