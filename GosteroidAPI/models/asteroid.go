package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Distance struct {
    Date     string  `json:"date" bson:"date"`
    Distance float64 `json:"distance" bson:"distance"`
}

type Asteroid struct {
    ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name          string             `json:"name" bson:"name"`
    Diameter      float64            `json:"diameter" bson:"diameter"`
    DiscoveryDate string             `json:"discovery_date" bson:"discovery_date"`
    Observations  string             `json:"observations,omitempty" bson:"observations,omitempty"`
    Distances     []Distance         `json:"distances,omitempty" bson:"distances,omitempty"`
}
