package getUrl

import (
	"github.com/AmadoMuerte/url-shortener/internal/lib/api/response"
	"github.com/AmadoMuerte/url-shortener/internal/lib/logger/logger"
	"github.com/AmadoMuerte/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

type UrlGetter interface {
	GetUrl(id int64) (sqlite.UrlData, error)
}

func New(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.url.getUrl"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var paramId = chi.URLParam(r, "id")
		reqId, err := strconv.ParseInt(paramId, 10, 64)
		if err != nil {
			log.Error("failed to convert id to int64", logger.Err(err))
			render.JSON(w, r, response.Error("invalid id format"))
			return
		}

		data, err := urlGetter.GetUrl(reqId)
		if err != nil {
			log.Error("url is not exist", logger.Err(err))
			render.JSON(w, r, response.Error("url is not exist"))

			return
		}

		response.UrlDateResponseOK(w, r, data)
	}
}
