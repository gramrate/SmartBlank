package game

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) handleAssignRole(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.AssignRoleRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateAssignRoleRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	game, err := h.gameService.AssignRole(context.Background(), &req)
	if err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("assign role failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "role assigned", game)
}
