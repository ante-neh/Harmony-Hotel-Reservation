package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const DatabaseName = "harmony_hotel"

func (m *MongoDb) GetCollection(collection string) (*mongo.Collection, error) {
	return m.Client.Database(DatabaseName).Collection(collection), nil
}
