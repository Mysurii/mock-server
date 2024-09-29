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

func (s *server) StartServer(port uint) error {
  
	handler := s.registerRoutes()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %v", err.Error()))
	}

	return nil
}

