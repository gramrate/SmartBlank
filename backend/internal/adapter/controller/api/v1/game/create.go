package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleCreate(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.CreateGameRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateCreateGameRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.Create(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("create game failed: %w", err).Error())
		return
	}
	if game != nil {
		h.bindLobbyID(ws, game.LobbyID)
	}
	h.wsUtils.SendSuccess(ws, "game created", game)
}
