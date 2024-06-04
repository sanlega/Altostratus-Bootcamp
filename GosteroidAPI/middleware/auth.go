package middleware

import (
	"GosteroidAPI/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

// JwtMiddleware returns the middleware for JWT authentication
func JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(utils.GetEnv("JWT_SECRET")),
	})
}

// GenerateJWT generates a JWT token for testing purposes
func GenerateJWT(c echo.Context) error {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "GosteroidAPI",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(utils.GetEnv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "could not generate token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
