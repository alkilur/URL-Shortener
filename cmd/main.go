package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// TODO: init config - cleanenv (минимализм, в отличии от Viper и Cobra)
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger - slog

	// TODO: init storage - sqlite3

	// TODO: init router - chi, chi render

	// TODO: run server

}
