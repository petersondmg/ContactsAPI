package contact

import (
	"capi/cmd/api/internal/middleware/auth"
	"capi/cmd/api/internal/middleware/clientservice"
	"capi/cmd/api/internal/middleware/log"
	"capi/domain/service"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

func Add(svc *service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		jwtClaims := auth.JWTFromContext(ctx)
		logger := log.FromContext(ctx).With(
			zap.Int("client_id", jwtClaims.ClientID),
			zap.String("client_name", jwtClaims.ClientName),
		)

		var payload ContactsPayload
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			logger.Warn("invalid payload")
			// invalid payload
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// get service for specific client
		clientSvc := clientservice.FromContext(ctx)

		contacts := payload.Entities()

		err := clientSvc.AddContacts(req.Context(), payload.Entities())
		switch {
		case errors.Is(err, service.ErrInvalidPhone), errors.Is(err, service.ErrInvalidName):
			// invalid payload - bad request
			logger.Warn("invalid contact", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return

		case err != nil:
			// unhandled error
			logger.Error("error saving contact", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger.Info("contacts created")

		if err = json.NewEncoder(w).Encode(NewContactListView(contacts)); err != nil {
			logger.Warn("error sending response", zap.Error(err))
		}
	}
}

func List(service *service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.FromContext(r.Context()).Warn("list method not implemented")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "method not implemented"}`))
	}
}
