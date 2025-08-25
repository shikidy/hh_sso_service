package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/shikidy/hh_sso_service/internal/store"
)

type HTTPApp struct {
	router *http.ServeMux
	server *http.Server
	logger *slog.Logger
	port   string
}

func New(store store.Store, logger *slog.Logger, port string) *HTTPApp {
	router := http.NewServeMux()

	authHandler := NewAuthHandler(store, logger)
	authHandler.Subscribe(router)

	return &HTTPApp{
		router: router,
		logger: logger,
		port:   port,
	}
}

func (a *HTTPApp) Run() error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", a.port),
		Handler: a.router,
	}
	a.logger.Info("Starting web applicationt at:" + a.port)
	a.server = &server
	return a.server.ListenAndServe()
}

func (a *HTTPApp) GracefulShutDown() error {
	a.logger.Info("Stopping web application")
	return a.server.Shutdown(context.Background())
}
