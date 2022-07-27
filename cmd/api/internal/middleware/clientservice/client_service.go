package clientservice

import (
	"capi/cmd/api/internal/middleware/auth"
	"capi/domain/service"
	"context"
	"net/http"
)

func Apply(svc *service.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtClaims := auth.JWTFromContext(r.Context())

			var clientService service.ClientService

			// For the time being using a switcw. We could use a register approach to dynamically
			// select a client
			switch jwtClaims.ClientName {
			case "macapa":
				clientService = svc.Macapa()

			case "varejao":
				clientService = svc.Varejao()
			}

			ctx := WithContext(r.Context(), clientService)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type contextKey struct{}

func WithContext(ctx context.Context, s service.ClientService) context.Context {
	return context.WithValue(ctx, contextKey{}, s)
}

func FromContext(ctx context.Context) service.ClientService {
	return ctx.Value(contextKey{}).(service.ClientService)
}
