package main

import (
	"fmt"
	_ "log/slog"
	config "url-shortener/internal"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi

	// TODO: init server
}
