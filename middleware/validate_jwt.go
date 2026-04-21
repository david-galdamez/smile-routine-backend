package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/david-galdamez/smile-routine-backend/utils"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "No autorizado")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Formato de token incorrecto")
			return
		}

		tokenstr := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			utils.RespondWithError(w, http.StatusInternalServerError, "JWT_SECRET no configurado")
			return
		}

		token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "Token inválido")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Token inválido")
			return
		}

		userID, ok := claims["userId"].(float64)
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Id inválido en token")
			return
		}

		authUser := utils.AuthUser{
			UserId: int(userID),
		}
		ctx := context.WithValue(r.Context(), "auth_user", authUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
