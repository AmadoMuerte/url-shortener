package update

import (
	"errors"
	"github.com/AmadoMuerte/url-shortener/internal/lib/api/response"
	"github.com/AmadoMuerte/url-shortener/internal/lib/logger"
	"github.com/AmadoMuerte/url-shortener/internal/storage"
	"github.com/AmadoMuerte/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"strconv"
)

type UrlUpdater interface {
	UpdateUrl(id int64, url, alias string) (sqlite.UrlData, error)
	CheckUrlExist(id int64) (int64, error)
}

type Request struct {
	Url   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

func New(log *slog.Logger, urlUpdater UrlUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.url.update"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())))

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", logger.Err(err))
			render.JSON(w, r, response.Error("failed to decode request"))

			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", logger.Err(err))
			render.JSON(w, r, response.Error("invalid request"))

			return
		}

		paramId := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(paramId, 10, 64)
		if err != nil {
			log.Error("failed to parse id", logger.Err(err))
			render.JSON(w, r, response.Error("failed to parse id"))

			return
		}
		if _, err := urlUpdater.CheckUrlExist(id); err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Error("url not found", logger.Err(err))
				render.JSON(w, r, response.Error("url not found"))
			}

			return
		}

		updatedUrl, err := urlUpdater.UpdateUrl(id, req.Url, req.Alias)
		if err != nil {
			log.Error("failed to update url", logger.Err(err))
			render.JSON(w, r, response.Error("failed to update url"))

			return
		}

		response.UrlDateResponseOK(w, r, updatedUrl)
	}
}
