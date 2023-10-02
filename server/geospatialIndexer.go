package main

import (
	"context"

	"github.com/SaiNageswarS/go-api-boot/odm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateGeospatialIndex(database string, collection string) {
	client := odm.GetClient()
	coll := client.Database(database).Collection(collection)
	//index database for geospatial queries
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "location", Value: "2dsphere"}},
	}
	_, err := coll.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}
}
