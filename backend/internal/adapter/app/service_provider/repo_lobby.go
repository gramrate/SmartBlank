package service_provider

import lobbyrepo "backend/internal/adapter/repo/mongo/lobby"

func (s *ServiceProvider) LobbyRepository() *lobbyrepo.Repo {
	if s.lobbyRepo == nil {
		s.lobbyRepo = lobbyrepo.NewRepo(s.MongoDB())
	}
	return s.lobbyRepo
}
