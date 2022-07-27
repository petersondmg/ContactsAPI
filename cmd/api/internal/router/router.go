package router

import (
	"capi/cmd/api/internal/handlers/contact"
	"capi/cmd/api/internal/middleware/auth"
	"capi/cmd/api/internal/middleware/clientservice"
	"capi/cmd/api/internal/middleware/log"
	"capi/domain/service"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func New(svc *service.Service, logger *zap.Logger, jwtSecret string) *mux.Router {
	r := mux.NewRouter()

	r.Use(auth.Apply(jwtSecret))
	r.Use(log.Apply(logger))
	r.Use(clientservice.Apply(svc))

	r.HandleFunc("/contact", contact.Add(svc)).Methods(http.MethodPost)
	r.HandleFunc("/contact", contact.List(svc)).Methods(http.MethodGet)

	return r
}
