package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	srv, err := server.New(logger)
	if err != nil {
		logger.Fatal("Ошибка создания сервера:", err)
	}

	srv.RegisterHandlers()

	if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Ошибка запуска сервера:", err)
	}
}
