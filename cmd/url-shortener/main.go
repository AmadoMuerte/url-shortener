package main

import (
	config "github.com/AmadoMuerte/url-shortener/internal"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/auth"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/redirect"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/getAllURL"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/remove"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/save"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/middleware/mwLogger"
	"github.com/AmadoMuerte/url-shortener/internal/lib/logger/logSlog"
	"github.com/AmadoMuerte/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log/slog"
	"net/http"
	"os"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Пост запрос сработает только под авторизацией
	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			cfg.HTTPServer.User: cfg.HTTPServer.Password,
		}))

		r.Get("/all", getAllURL.New(log, storage, cfg.Address))
		r.Post("/", save.New(log, storage))
		r.Put("/{id}", remove.New(log, storage))
	})

	router.Post("/auth/login", auth.New(log, cfg.HTTPServer.User, cfg.HTTPServer.Password))

	router.Get("/{alias}", redirect.New(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")
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
		// TODO: init mwLogger: logSlog
	default:
		// TODO: init mwLogger: logSlog
	}
	return log
}
