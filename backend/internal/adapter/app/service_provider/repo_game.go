package service_provider

import gamerepo "backend/internal/adapter/repo/mongo/game"

func (s *ServiceProvider) GameRepository() *gamerepo.GameRepository {
	if s.gameRepo == nil {
		s.gameRepo = gamerepo.NewGameRepository(s.MongoDB())
	}
	return s.gameRepo
}
