package response

import (
	"github.com/AmadoMuerte/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

type UrlDateResponse struct {
	Response
	Id    int64  `json:"id"`
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

func UrlDateResponseOK(w http.ResponseWriter, r *http.Request, urlData sqlite.UrlData) {
	render.JSON(w, r, UrlDateResponse{
		Response: OK(),
		Id:       urlData.Id,
		Url:      urlData.Url,
		Alias:    urlData.Alias,
	})
}
