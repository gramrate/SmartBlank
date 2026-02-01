package lobby

import (
	"backend/internal/adapter/controller/validator"
	"backend/internal/domain/dto"
	"context"

	"backend/internal/adapter/controller/api/utils/websocket"

	"github.com/labstack/echo/v4"
)

type lobbyService interface {
	Close(ctx context.Context, req *dto.CloseLobbyRequest) (*dto.CloseLobbyResponse, error)
	Create(ctx context.Context, req *dto.CreateLobbyRequest) (*dto.CreateLobbyResponse, error)
}

type Handler struct {
	lobbyService lobbyService
	wsUtils      *websocket.WebSocket
	validator    *validator.Validator
}

func NewHandler(lobbyService lobbyService, wsUtils *websocket.WebSocket, validator *validator.Validator) *Handler {
	return &Handler{
		lobbyService: lobbyService,
		wsUtils:      wsUtils,
		validator:    validator,
	}
}

func (h *Handler) Setup(router *echo.Group) {
	router.POST("/lobby/create", h.Create)
	router.GET("/ws/lobby", h.HandleWebSocket)
}
