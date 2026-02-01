package game

import "github.com/labstack/echo/v4"

func (h *Handler) Setup(router *echo.Group) {
	router.GET("/ws/game", h.HandleWebSocket)
}
