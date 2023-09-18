package odm

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connection *mongo.Client = nil

func newMongoConn() *mongo.Client {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		log.Fatal("mongouri is empty")
	}

	mongoOpts := options.Client().ApplyURI(mongoUri)
	//mongoOpts.TLSConfig.MinVersion = tls.VersionTLS12
	// make sure to install ca-certs in docker image.
	// mongoOpts.TLSConfig.InsecureSkipVerify = true

	client, err := mongo.Connect(context.TODO(), mongoOpts)
	if err != nil {
		panic(err)
	}

	return client
}

func GetClient() *mongo.Client {
	if connection != nil {
		return connection
	}

	connection = newMongoConn()
	return connection
}
