package lobby

import (
	"backend/internal/domain/dto"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

func (h *Handler) Close(ws *websocket.Conn, msg *dto.WSMessage) {
	var req dto.CloseLobbyRequest
	if err := h.wsUtils.UnmarshalData(msg.Data, &req); err != nil {
		h.wsUtils.SendError(ws, "invalid data format")
		return
	}
	if err := h.validator.ValidateCloseLobbyRequest(&req); err != nil {
		h.wsUtils.SendError(ws, err.Error())
		return
	}
	if _, err := h.lobbyService.Close(context.Background(), &req); err != nil {
		h.wsUtils.SendError(ws, fmt.Errorf("close lobby failed: %w", err).Error())
		return
	}
	h.wsUtils.SendSuccess(ws, "game closed", nil)
}
