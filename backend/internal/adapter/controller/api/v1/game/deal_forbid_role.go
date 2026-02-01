package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleSetForbiddenRole(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.ForbiddenRoleRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateForbiddenRoleRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.SetForbiddenRole(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("set forbidden role failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "forbidden roles updated", game)
}
