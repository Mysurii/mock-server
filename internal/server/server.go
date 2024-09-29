package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mysurii/mock-server/internal/config"
	"github.com/mysurii/mock-server/internal/models"
)

type server struct {
	api models.API
}

func New(jsonPath string) *server {
	api, err  := config.LoadApiFile(jsonPath)
	if err != nil {
		panic(err)
	}

	return &server{
		api: api,
	}
}

func (s *server) StartServer() error {
  
	handler := s.registerRoutes()

	chainedHandler := LoggingMiddleware(handler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.api.Port),
		Handler:      chainedHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Mock-server started on port %d\n\n", s.api.Port)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %v", err.Error()))
	}

	return nil
}

