package middleware

import (
	"InkaTry/warehouse-storage-be/internal/pkg/config"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func Middleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
			if len(authHeader) != 2 {
				fmt.Println("Malformed token")
				responder.ResponseError(w, errs.ErrAuth)
			} else {
				jwtToken := authHeader[1]
				token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					SECRETKEY := cfg.JWTSecret
					return []byte(SECRETKEY), nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.Background()
					ctx = context.WithValue(ctx, "email", claims["email"])
					ctx = context.WithValue(ctx, "admin", claims["admin"])

					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					fmt.Println(err)
					responder.ResponseError(w, errs.ErrAuth)

				}
			}
		})
	}
}
