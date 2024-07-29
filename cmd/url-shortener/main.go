package main

import (
	config "github.com/AmadoMuerte/url-shortener/internal"
	httpserver "github.com/AmadoMuerte/url-shortener/internal/http-server"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/router"
	"github.com/AmadoMuerte/url-shortener/internal/lib/logger/logger"
	"github.com/AmadoMuerte/url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", logger.Err(err))
		os.Exit(1)
	}
	r := router.New(cfg, log, storage)
	httpserver.New(cfg, log, r)
}
