package router

import (
	config "github.com/AmadoMuerte/url-shortener/internal"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/auth"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/redirect"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/getAllURL"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/getUrl"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/remove"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/save"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/handlers/url/update"
	"github.com/AmadoMuerte/url-shortener/internal/http-server/middleware/mwLogger"
	"github.com/AmadoMuerte/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log/slog"
	"net/http"
)

func New(cfg *config.Config, log *slog.Logger, storage *sqlite.Storage) http.Handler {
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

	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("url.ts-shortener", map[string]string{
			cfg.HTTPServer.User: cfg.HTTPServer.Password,
		}))
		r.Get("/all", getAllURL.New(log, storage, cfg.Address))
		r.Get("/{id}", getUrl.New(log, storage))
		r.Post("/", save.New(log, storage))
		r.Put("/{id}", update.New(log, storage))
		r.Delete("/{id}", remove.New(log, storage))
	})

	router.Post("/auth/login", auth.New(log, cfg.HTTPServer.User, cfg.HTTPServer.Password))
	router.Get("/{alias}", redirect.New(log, storage))

	return router
}
