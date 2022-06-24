package handler

import (
	"net/http"

	"github.com/MIHAIL33/Ponylab-Go/pkg/service"
)

//Handler - type of handle
type Handler struct {
	service *service.Service
}

//NewHandler - constructor
func NewHandler(service *service.Service) *Handler {
	handler := &Handler{service: service}
	handler.initRoutes()
	return handler
}

func (h *Handler) initRoutes() {

	http.Handle("/", http.HandlerFunc(h.getBasePage))
	http.Handle("/ws", http.HandlerFunc(h.wsHandler))

}