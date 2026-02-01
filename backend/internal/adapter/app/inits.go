package app

import (
	"backend/internal/adapter/app/service_provider"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func (a *App) initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./backend")
	viper.AddConfigPath("../backend")

	if err := viper.ReadInConfig(); err != nil {
		if _, statErr := os.Stat("config.yaml"); statErr == nil {
			return err
		}
		if _, statErr := os.Stat("./backend/config.yaml"); statErr == nil {
			return err
		}
		if _, statErr := os.Stat("../backend/config.yaml"); statErr == nil {
			return err
		}
		return err
	}
	return nil
}

func (a *App) initServiceProvider() error {
	a.ServiceProvider = service_provider.New()
	return nil
}

// initHTTPServer initializes the Echo server
func (a *App) initHTTPServer() error {
	e := echo.New()
	a.Server = e
	return nil
}
