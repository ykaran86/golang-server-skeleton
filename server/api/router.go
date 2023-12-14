package api

import (
	"github.com/gorilla/mux"
	"golang-server-skeleton/server/handlers/ping"
	"net/http"
)

type handlers struct {
	pingHandler *ping.Handler
}

func createHandlers() *handlers {
	pingHandler := &ping.Handler{}

	return &handlers{
		pingHandler: pingHandler,
	}
}

func NewRouter() http.Handler {
	handlers := createHandlers()

	router := mux.NewRouter()

	router.HandleFunc("/ping", handlers.pingHandler.HandlePing).Methods(http.MethodGet)

	return router
}
