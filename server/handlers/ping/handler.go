package ping

import (
	"google.golang.org/appengine/log"
	"net/http"
)

type Handler struct{}

func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Errorf(r.Context(), "Failed Handle Ping, err: %v", err)
	}
}
