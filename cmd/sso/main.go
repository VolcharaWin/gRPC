package main

import (
	"fmt"
	"internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: инциализировать объект конфига

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}