package handler

import (
	"m2m/internal/shared"
	"net/http"
	"strings"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return new(HealthHandler)
}

func (th *HealthHandler) GetRoutes() shared.TypeRouters {
	return shared.TypeRouters{
		"/health": th.handler,
	}
}

func (th *HealthHandler) handler(w http.ResponseWriter, r *http.Request) {
	switch strings.ToUpper(r.Method) {
	case "GET":
		th.getHealths(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (th *HealthHandler) getHealths(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("A live"))
}
