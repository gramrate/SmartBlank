package lobby

import (
	"backend/internal/domain/dto"
	"encoding/json"
	"log"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HandleWebSocket(c echo.Context) error {
	ws, err := h.wsUtils.Upgrader().Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Читаем сообщение от клиента
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Ошибка чтения: %v", err)
			break
		}

		var wsMsg dto.WSMessage
		if err := json.Unmarshal(msg, &wsMsg); err != nil {
			h.wsUtils.SendError(ws, "Неверный формат сообщения")
			continue
		}

		// Маршрутизация сообщений
		switch wsMsg.Type {
		case "close_lobby":
			h.Close(ws, &wsMsg)
		default:
			h.wsUtils.SendError(ws, "Неизвестный тип сообщения")
		}
	}

	return nil
}
