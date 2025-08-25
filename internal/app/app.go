package app

import (
	"log/slog"

	"github.com/shikidy/hh_sso_service/internal/app/web"
	"github.com/shikidy/hh_sso_service/internal/config"
	"github.com/shikidy/hh_sso_service/internal/store"
)

type App struct {
	store  store.Store
	logger *slog.Logger
	config *config.Config

	Webapp *web.HTTPApp
}

func New(
	store store.Store,
	logger *slog.Logger,
	config *config.Config,
) *App {
	return &App{
		store:  store,
		logger: logger,
		config: config,
		Webapp: web.New(store, logger, config.HTTP.Port),
	}
}
