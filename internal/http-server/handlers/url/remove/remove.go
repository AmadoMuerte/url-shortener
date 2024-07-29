package remove

import (
	"github.com/AmadoMuerte/url-shortener/internal/lib/api/response"
	"github.com/AmadoMuerte/url-shortener/internal/lib/logger/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

type UrlRemover interface {
	RemoveUrl(int64) (int64, error)
}

type Response struct {
	response.Response
	Id int64 `json:"id,omitempty"`
}

func New(log *slog.Logger, urlRemover UrlRemover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.url.remove"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		reqIdStr := chi.URLParam(r, "id")
		reqId, err := strconv.ParseInt(reqIdStr, 10, 64)
		if err != nil {
			log.Error("failed to convert id to int64", logger.Err(err))
			render.JSON(w, r, response.Error("invalid id format"))
			return
		}

		id, err := urlRemover.RemoveUrl(reqId)
		if err != nil {
			log.Error("url is not exist", logger.Err(err))
			render.JSON(w, r, response.Error("url is not exist"))
			return
		}

		log.Info("url has been deleted", slog.Int64("id", id))

		responseOK(w, r, id)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, id int64) {
	render.JSON(w, r, Response{
		Response: response.OK(),
		Id:       id,
	})
}
