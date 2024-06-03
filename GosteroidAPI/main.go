package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/your_project_name/handlers"
    "github.com/your_project_name/middleware"
    "github.com/your_project_name/models"
    "github.com/your_project_name/utils"
)

func main() {
    utils.LoadEnv()
    models.InitDatabase()

    r := mux.NewRouter()

    api := r.PathPrefix("/api/v1").Subrouter()

    api.HandleFunc("/asteroides", handlers.CreateAsteroid).Methods("POST")
    api.HandleFunc("/asteroides", handlers.GetAsteroids).Methods("GET")
    api.HandleFunc("/asteroides/{id}", handlers.GetAsteroidByID).Methods("GET")
    api.HandleFunc("/asteroides/{id}", handlers.UpdateAsteroid).Methods("PATCH")
    api.HandleFunc("/asteroides/{id}", handlers.DeleteAsteroid).Methods("DELETE")

    // JWT protected routes
    api.Use(middleware.JwtVerify)

    log.Fatal(http.ListenAndServe(":8080", r))
}
