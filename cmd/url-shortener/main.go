package main

import (
	"log/slog"
	"os"
	config "url-shortener/internal"
	"url-shortener/internal/lib/logger/logSlog"
	"url-shortener/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", logSlog.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init router: chi

	// TODO: init server
}

func setupLogger(env string) *slog.Logger {
	const (
		envLocal = "local"
		envDev   = "dev"
		envProd  = "prod"
	)

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
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		// TODO: init logger: logSlog
	default:
		// TODO: init logger: logSlog
	}

	return log
}
