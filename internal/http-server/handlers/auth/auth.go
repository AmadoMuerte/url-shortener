package auth

import (
	"github.com/AmadoMuerte/url-shortener/internal/lib/api/response"
	"github.com/AmadoMuerte/url-shortener/internal/lib/logger"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type Request struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Response struct {
	response.Response
}

func New(log *slog.Logger, cfgUser, cfgPass string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "handlers.auth.checkAuth"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

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

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", logger.Err(err))
			render.JSON(w, r, response.Error("invalid request"))

			return
		}

		login := req.Login
		password := req.Password
		if login != cfgUser || password != cfgPass {
			log.Error("wrong data")
			render.JSON(w, r, response.Error("wrong data"))

			return
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
		})
	}
}
