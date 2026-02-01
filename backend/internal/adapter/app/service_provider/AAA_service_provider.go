package service_provider

import (
	"backend/internal/adapter/controller/validator"
	wsUtils "backend/internal/adapter/controller/api/utils/websocket"
	gamerepo "backend/internal/adapter/repo/mongo/game"
	lobbyrepo "backend/internal/adapter/repo/mongo/lobby"
	"backend/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceProvider struct {
	serverConfig serverConfig
	loggerConfig loggerConfig
	mongoConfig  mongoConfig
	lobbyService lobbyService
	gameService  gameService
	lobbyRepo    *lobbyrepo.Repo
	validator    *validator.Validator

	mongoDB *mongo.Database

	logger   *logger.Logger
	wsUtils  *wsUtils.WebSocket
	gameRepo *gamerepo.GameRepository
}

func New() *ServiceProvider {
	return &ServiceProvider{}
}
