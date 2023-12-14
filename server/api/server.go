package api

import (
	"context"
	"fmt"
	"golang-server-skeleton/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	cfg := config.GetServerConfig()
	return &Server{
		server: &http.Server{
			Handler:      NewRouter(),
			Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.Port),
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		},
	}
}

func (s *Server) Start() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func(s *http.Server) {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server, err: %v", err)
		}
	}(s.server)

	<-sig

	err := s.server.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("Error shutting down server, err: %v", err)
	}
}
