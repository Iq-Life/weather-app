package main

import (
	"fmt"

	"github.com/Iq-Life/weather-app/internal/config"
)

func main() {
	// TODO: init config cleanenv
	// TODO: init logger log/slog
	// TODO: init storage sqlite
	// TODO: init router chi, render
	// TODO: init server

	cfg := config.MustLoad()

	fmt.Println("go world", cfg)
}
