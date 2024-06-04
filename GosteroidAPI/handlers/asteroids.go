package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"

	"GosteroidAPI/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection() *mongo.Collection {
	return models.DB.Collection("asteroids")
}

func GetAsteroids(c echo.Context) error {
	collection := getCollection()

	// Get pagination parameters
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	// Default values for pagination
	page := 1
	limit := 10

	// Parse page parameter
	if pageParam != "" {
		parsedPage, err := strconv.Atoi(pageParam)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Parse limit parameter
	if limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Calculate skip
	skip := (page - 1) * limit

	// MongoDB options for pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))

	cur, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch asteroids"})
	}
	defer cur.Close(context.Background())

	var asteroids []models.Asteroid
	for cur.Next(context.Background()) {
		var asteroid models.Asteroid
		if err := cur.Decode(&asteroid); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error decoding asteroid"})
		}
		asteroids = append(asteroids, asteroid)
	}

	if err := cur.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error iterating through cursor"})
	}

	return c.JSON(http.StatusOK, asteroids)
}

func CreateAsteroid(c echo.Context) error {
	collection := getCollection()
	var asteroid models.Asteroid
	if err := c.Bind(&asteroid); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	asteroid.ID = primitive.NewObjectID()

	result, err := collection.InsertOne(context.Background(), asteroid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create asteroid"})
	}
	return c.JSON(http.StatusCreated, result)
}

func GetAsteroidByID(c echo.Context) error {
	collection := getCollection()
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	var asteroid models.Asteroid
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&asteroid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch asteroid"})
	}
	return c.JSON(http.StatusOK, asteroid)
}

func UpdateAsteroid(c echo.Context) error {
	collection := getCollection()
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	var asteroid models.Asteroid
	if err := c.Bind(&asteroid); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	update := bson.M{"$set": asteroid}
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update asteroid"})
	}
	return c.JSON(http.StatusOK, asteroid)
}

func DeleteAsteroid(c echo.Context) error {
	collection := getCollection()
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete asteroid"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Asteroide eliminado exitosamente"})
}
