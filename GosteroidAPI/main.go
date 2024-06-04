package main

import (
	"net/http"

	"GosteroidAPI/handlers"
	"GosteroidAPI/middleware"
	"GosteroidAPI/models"
	"GosteroidAPI/utils"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.LoadEnv()
	models.InitDatabase()

	e := echo.New()

	// Middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	// JWT generation route for testing
	e.POST("/login", middleware.GenerateJWT)

	// JWT protected routes
	api := e.Group("/api/v1")
	api.Use(middleware.JwtMiddleware())

	api.POST("/asteroides", handlers.CreateAsteroid)
	api.GET("/asteroides", handlers.GetAsteroids)
	api.GET("/asteroides/:id", handlers.GetAsteroidByID)
	api.PATCH("/asteroides/:id", handlers.UpdateAsteroid)
	api.DELETE("/asteroides/:id", handlers.DeleteAsteroid)

	// Custom 404 handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		c.JSON(code, map[string]string{"message": "Not Found"})
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
