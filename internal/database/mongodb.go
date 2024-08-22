package database

import "go.mongodb.org/mongo-driver/mongo"

type MongoDb struct {
	Client *mongo.Client
}