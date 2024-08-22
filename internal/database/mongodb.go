package database

import (
	"context"
	"github.com/ante-neh/Harmony-Hotel-Reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserStore interface{
	GetUser(context.Context, string)(*types.User, error)
	CreateUser(context.Context, string, string, string)(types.User, error)
}
type MongoDb struct {
	Client *mongo.Client
}


func (m *MongoDb) CreateUser(ctx context.Context, userName, email, password string)(*types.User, error){
	collection, err := m.GetCollection("users")
	if err != nil{
		return &types.User{}, err
	}
	// user := types.User{}
	_, err = collection.InsertOne(ctx, bson.M{})

	if err != nil{
		return &types.User{}, err
	}

	return &types.User{}, nil 
} 

func (m *MongoDb) GetUser(ctx context.Context, id string)(*types.User, error){
	collection, err := m.GetCollection("users")
	user := types.User{} 

	if err != nil{
		return &types.User{}, err
	}

	objectId, err := ToBsonObject(id)

	if err != nil{
		return &types.User{}, err
	}
	
	err = collection.FindOne(ctx, bson.M{"_id":objectId}).Decode(&user)

	if err != nil{
		return &types.User{}, err
	}

	return &user, nil 
}


func (m *MongoDb) GetUsers(ctx context.Context)([]*types.User, error){
	collection, err := m.GetCollection("users")
	if err != nil{
		return []*types.User{}, err
	}
	result, err := collection.Find(ctx, bson.M{}) 

	if err != nil{
		return []*types.User{}, err
	}

	users := []*types.User{} 

	err = result.Decode(&users) 

	if err != nil{
		return []*types.User{}, err
	}

	return users, nil 
}
