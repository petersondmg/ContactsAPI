package log

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

func Apply(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// add other useful info
			l := logger.With(zap.String("url", r.URL.String()))

			// insert logger into
			ctx := WithContext(r.Context(), l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type contextKey struct{}

func WithContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}

func FromContext(ctx context.Context) *zap.Logger {
	return ctx.Value(contextKey{}).(*zap.Logger)
}
