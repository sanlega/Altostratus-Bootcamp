package middleware

import (
    "context"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/your_project_name/utils"
)

type key int

const (
    UserKey key = iota
)

func JwtVerify(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, http.ErrAbortHandler
            }
            return []byte(utils.GetEnv("JWT_SECRET")), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            ctx := context.WithValue(r.Context(), UserKey, claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            http.Error(w, err.Error(), http.StatusUnauthorized)
        }
    })
}
