package database

import (
	"github.com/ante-neh/Harmony-Hotel-Reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
)


type Database interface{
	CreateUser(userName, email, password string)(types.User, error)
}
type MongoDb struct {
	Client *mongo.Client
}


func (m *MongoDb) CreateUser(userName, email, password string)(types.User, error){
	return types.User{}, nil 
} 

func (m *MongoDb) GetUser(id string)(types.User, error){
	return types.User{}, nil 
}