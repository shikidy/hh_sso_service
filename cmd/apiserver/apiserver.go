package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/shikidy/hh_sso_service/internal/app"
	"github.com/shikidy/hh_sso_service/internal/config"
	"github.com/shikidy/hh_sso_service/internal/store/sqlstore"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/local.yaml", "set application config")
}

func main() {
	flag.Parse()

	cfg := config.MustLoadConfig(configPath)
	logger := setupLogger(cfg.Env)

	db, err := newDB(cfg.StoragePath)
	if err != nil {
		panic("Error on loading database: " + err.Error())
	}
	store := sqlstore.New(db)
	application := app.New(store, logger, cfg)

	go application.Webapp.Run()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	application.Webapp.GracefulShutDown()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func newDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
