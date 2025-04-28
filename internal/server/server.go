package server

import (
	"log"
	"net/http"
	"time"

	handlers "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server представляет собой структуру HTTP-сервера
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// New создает новый сервер
func New(logger *log.Logger) (*Server, error) {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: srv,
	}, nil
}

// RegisterHandlers регистрирует хэндлеры в маршрутизаторе
func (s *Server) RegisterHandlers() {
	mux := s.Server.Handler.(*http.ServeMux)

	// Регистрируем хэндлер только для точного пути '/'
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler(s.Logger))
}
