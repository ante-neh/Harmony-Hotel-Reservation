package database

import (
	"context"
	"github.com/ante-neh/Harmony-Hotel-Reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserStore interface{
	GetUsers(context.Context)([]*types.User, error)
	DeleteUser(context.Context, string)(string, error)
	GetUser(context.Context, string)(*types.User, error)
	CreateUser(context.Context, *types.User)(*types.User, error)
	UpdateUser(context.Context, primitive.ObjectID)(*types.User, error)
}
type MongoDb struct {
	Client *mongo.Client
}


func (m *MongoDb) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	collection, err := m.GetCollection("users")
	if err != nil {
		return nil, err
	}

	user.ID = primitive.NewObjectID()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m *MongoDb) GetUser(ctx context.Context, id primitive.ObjectID)(*types.User, error){
	collection, err := m.GetCollection("users")
	user := types.User{} 

	if err != nil{
		return &types.User{}, err
	}
	
	
	err = collection.FindOne(ctx, bson.M{"_id":id}).Decode(&user)

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

	err = result.All(ctx, &users) 

	if err != nil{
		return []*types.User{}, nil
	}

	return users, nil 
}


func (m *MongoDb) DeleteUser(ctx context.Context, id primitive.ObjectID) (string, error){
	collection, err := m.GetCollection("users") 

	if err != nil{
		return "table doesn't exist", err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id":id})
	
	if err != nil{
		return "", err
	}
	return "User Deleted", nil 
}


func (m *MongoDb) UpdateUser(ctx context.Context, id primitive.ObjectID, user types.UserRequest)(*types.User, error){
	collection, err := m.GetCollection("users")
	if err != nil{
		return &types.User{}, err
	}

	update := bson.M{
		"$set": user,
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil{
		return nil, err
	}

	userResponse := types.User{
		ID:user.ID,
		Name:user.Name,
		Email:user.Email,
	}
	return &userResponse, nil 
}