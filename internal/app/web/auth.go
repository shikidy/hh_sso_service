package web

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/shikidy/hh_sso_service/internal/store"
)

type AuthRouter struct {
	router *http.ServeMux
	store  store.Store
	logger *slog.Logger
}

func NewAuthHandler(store store.Store, logger *slog.Logger) *AuthRouter {
	router := http.NewServeMux()
	autRouter := &AuthRouter{
		router: router,
		store:  store,
		logger: logger,
	}
	autRouter.RegisterHandlers()
	return autRouter
}

func (r *AuthRouter) Subscribe(handler *http.ServeMux) {
	handler.Handle("/auth/", http.StripPrefix("/auth", r.router))
}

func (r *AuthRouter) RegisterHandlers() {
	r.logger.Debug("Registering auth handlers")
	r.router.HandleFunc("POST /register", r.HandleRegister())
}

func (r *AuthRouter) HandleRegister() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hi!")
	}
}
