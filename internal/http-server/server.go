package http_server

import (
	config "github.com/AmadoMuerte/url-shortener/internal"
	"log/slog"
	"net/http"
	"os"
)

func New(cfg *config.Config, log *slog.Logger, router http.Handler) {
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	log.Info("starting server", slog.String("address", cfg.Address))
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
		os.Exit(1)
	}
	log.Error("server stopped")
}
