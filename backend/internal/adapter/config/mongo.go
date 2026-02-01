package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// MongoConfig holds MongoDB configuration
type MongoConfig struct {
	// Basic connection
	host       string
	port       int
	database   string
	username   string
	password   string
	authSource string
}

// NewMongoConfig initializes a new MongoDB configuration from environment variables.
func NewMongoConfig() (*MongoConfig, error) {
	return &MongoConfig{
		host:       viper.GetString("mongodb.host"),
		port:       viper.GetInt("mongodb.port"),
		database:   viper.GetString("mongodb.database"),
		username:   viper.GetString("mongodb.username"),
		password:   viper.GetString("mongodb.password"),
		authSource: viper.GetString("mongodb.auth_source"),
	}, nil
}

// Host returns MongoDB host
func (cfg *MongoConfig) Host() string {
	return cfg.host
}

// Port returns MongoDB port
func (cfg *MongoConfig) Port() int {
	return cfg.port
}

// Database returns default database name
func (cfg *MongoConfig) Database() string {
	return cfg.database
}

// Username returns MongoDB username
func (cfg *MongoConfig) Username() string {
	return cfg.username
}

// Password returns MongoDB password
func (cfg *MongoConfig) Password() string {
	return cfg.password
}

// AuthSource returns MongoDB authentication source
func (cfg *MongoConfig) AuthSource() string {
	if cfg.authSource == "" {
		return "admin"
	}
	return cfg.authSource
}

// UriAddr возвращает полный URI для подключения к MongoDB
func (cfg *MongoConfig) UriAddr() string {
	// Если есть логин и пароль - добавляем их в URI
	if cfg.username != "" && cfg.password != "" {
		// Формат: mongodb://username:password@host:port/database?authSource=...
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			cfg.username,
			cfg.password,
			cfg.host,
			cfg.port,
			cfg.database)

		// Добавляем authSource если он указан
		authSource := cfg.AuthSource() // Используем метод для получения дефолтного значения
		if authSource != "" {
			uri += fmt.Sprintf("?authSource=%s", authSource)
		}

		return uri
	}

	// Если аутентификации нет - простой URI
	return fmt.Sprintf("mongodb://%s:%d/%s", cfg.host, cfg.port, cfg.database)
}
