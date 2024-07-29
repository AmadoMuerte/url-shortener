package getAllURL

import (
	"github.com/AmadoMuerte/url-shortener/internal/lib/api/response"
	"github.com/AmadoMuerte/url-shortener/internal/lib/logger"
	"github.com/AmadoMuerte/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type URLGetter interface {
	GetAllAlias() ([]sqlite.UrlInfo, error)
}

type Response struct {
	response.Response
	Data    []sqlite.UrlInfo `json:"data"`
	Address string           `json:"address"`
}

func New(log *slog.Logger, URLGetter URLGetter, address string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.getAllURL"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		data, err := URLGetter.GetAllAlias()
		if err != nil {
			log.Error("error getting aliases", logger.Err(err))
			render.JSON(w, r, response.Error("error getting aliases"))
			return
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
			Data:     data,
			Address:  address,
		})
	}
}
