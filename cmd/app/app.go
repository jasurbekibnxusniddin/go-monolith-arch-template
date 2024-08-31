package app

import (
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/config"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/api"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/pkgs/log"
	"github.com/jasurbekibnxusniddin/go-monolith-arch-template/internal/service"
)

func StartApp(cfg config.Config) {

	log := log.NewLogger(cfg.GeneralConfig)

	service := service.NewService(cfg, log)
	defer service.CleanUP()

	engine := api.Api(
		api.Options{
			Service: service,
			Log:     log,
		},
	)

	engine.Run(cfg.GeneralConfig.HTTPPort)
}
