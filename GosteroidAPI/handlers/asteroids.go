package handlers

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/your_project_name/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = models.DB.Collection("asteroids")

func CreateAsteroid(w http.ResponseWriter, r *http.Request) {
    var asteroid models.Asteroid
    json.NewDecoder(r.Body).Decode(&asteroid)
    asteroid.ID = primitive.NewObjectID()

    result, err := collection.InsertOne(context.Background(), asteroid)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}

func GetAsteroids(w http.ResponseWriter, r *http.Request) {
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    var asteroids []models.Asteroid
    for cur.Next(context.Background()) {
        var asteroid models.Asteroid
        cur.Decode(&asteroid)
        asteroids = append(asteroids, asteroid)
    }
    json.NewEncoder(w).Encode(asteroids)
}

func GetAsteroidByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var asteroid models.Asteroid
    err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&asteroid)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(asteroid)
}

func UpdateAsteroid(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var asteroid models.Asteroid
    json.NewDecoder(r.Body).Decode(&asteroid)

    update := bson.M{"$set": asteroid}
    _, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(asteroid)
}

func DeleteAsteroid(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    _, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"message": "Asteroide eliminado exitosamente"})
}
