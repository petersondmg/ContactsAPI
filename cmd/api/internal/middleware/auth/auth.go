package auth

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	ClientID   int    `json:"client_id"`
	ClientName string `json:"client_name"`
	jwt.StandardClaims
}

func Apply(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := JWTClaims{}
			jwtHeader := r.Header.Get("Authorization")

			token, err := jwt.ParseWithClaims(jwtHeader, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// insert key on request context
			ctx := WithContext(r.Context(), claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type contextKey struct{}

func WithContext(ctx context.Context, k JWTClaims) context.Context {
	return context.WithValue(ctx, contextKey{}, k)
}

func JWTFromContext(ctx context.Context) JWTClaims {
	return ctx.Value(contextKey{}).(JWTClaims)
}

/*
func init() {
		for i, name := range []string{"macapa", "varejao"} {
			claims := &JWTClaims{
				ClientID:   i + 1,
				ClientName: name,
				StandardClaims: jwt.StandardClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, _ := token.SignedString(jwtSecret)
			fmt.Println(name, tokenString)
		}
}
*/
