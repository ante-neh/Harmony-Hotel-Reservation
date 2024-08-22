package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DatabaseName = "harmony_hotel"

func (m *MongoDb) GetCollection(collection string) (*mongo.Collection, error) {
	return m.Client.Database(DatabaseName).Collection(collection), nil
}

func ToBsonObject(id string) (primitive.ObjectID, error){
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil{
		return primitive.ObjectID{}, err
	}

	return objectId, nil 
}