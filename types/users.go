package types

import (
	"regexp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
const(
	emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Name     string    `bson:"name" json:"name"`
	Email    string    `bson:"email" json:"email"`
	Password string    `bson:"password" json:"-"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequest) ValidateUser()map[string]string{
	errors := make(map[string]string) 
	if len(u.Name) < 2{
		errors["name"] = "user name should be more than two characters"
	}

	if len(u.Password) < 8{
		errors["password"] = "Password should have at least 8 characters"
	}

	if !isValid(u.Email){
		errors["email"] = "Email is invalid"
	}

	return errors
	
}

func isValid(e string) bool{
	emailRegx := regexp.MustCompile(emailRegexPattern)
	return emailRegx.MatchString(e)
}