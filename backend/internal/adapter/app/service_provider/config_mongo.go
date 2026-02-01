package service_provider

import (
	"backend/internal/adapter/config"
	"fmt"
)

type mongoConfig interface {
	Host() string
	Port() int
	Database() string
	Username() string
	Password() string
	AuthSource() string
	UriAddr() string
}

func (s *ServiceProvider) MongoConfig() mongoConfig {
	if s.mongoConfig == nil {
		cfg, err := config.NewMongoConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get mongo config: %w", err))
		}

		s.mongoConfig = cfg
	}

	return s.mongoConfig

}
