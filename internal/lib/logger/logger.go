package logger

import (
	"github.com/MatusOllah/slogcolor"
	"log/slog"
	"os"
)

func New(mode string) *slog.Logger {
	//TODO create logger based on mode
	var log *slog.Logger

	log = slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions))
	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
