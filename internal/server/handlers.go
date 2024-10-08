package server

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/ante-neh/Harmony-Hotel-Reservation/types"
	"github.com/ante-neh/Harmony-Hotel-Reservation/util"
)

func (s *Server) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	util.ResponseWithJson(w, 200, map[string]string{"message": "server is up"})
}


func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	params := types.UserRequest{} 
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil{
		util.ResponseWithError(w, 400, "Bad request")
		return 
	}

	//validate user data
	if errors := params.ValidateUser(); len(errors) > 0{
		util.ResponseWithJson(w, 400, errors)
		return 
	}

	//hash the password
	hashedPassword, err := util.Encrypt(params.Password, s.Bcrypt) 
	if err != nil{
		util.ResponseWithError(w, 400, "Bad request")
		return 
	}


	user := &types.User{
		Name:params.Name,
		Email: params.Email,
		Password: hashedPassword,
	}


	ctx := context.Background()
	userData, err := s.DB.CreateUser(ctx, user)

	if err != nil{
		util.ResponseWithError(w, 400, "Couldn't create user")
	}

	util.ResponseWithJson(w, 201, userData)
}


func (s *Server) HandleGetUser(w http.ResponseWriter, r *http.Request, user types.User){
	ctx := context.Background() 
	userData, err := s.DB.GetUser(ctx, user.ID)

	if err != nil{
		s.ErrorLogger.Println(err)
		util.ResponseWithError(w, 400, "Couldn't get user")
		return 
	}

	util.ResponseWithJson(w, 200, userData)
	
}


func (s *Server) HandleGetUsers(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	users, err := s.DB.GetUsers(ctx) 
	
	if err != nil{
		s.ErrorLogger.Println(err)
		util.ResponseWithError(w, 400, "Couldn't get users")
		return 
	}

	util.ResponseWithJson(w, 200, users)
}


func (s *Server) HandleDeleteUser(w http.ResponseWriter, r *http.Request, user types.User){
	ctx := context.Background()
	response, err := s.DB.DeleteUser(ctx, user.ID) 

	if err != nil {
		s.ErrorLogger.Println(err)
		util.ResponseWithError(w, 400, "User doesn't exist")
		return 
	}


	util.ResponseWithJson(w, 200, map[string]string{"message":response} )
}

func (s *Server) HandleUpdateUser(w http.ResponseWriter, r *http.Request, user types.User){
	params := types.UserRequest{} 
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil{
		util.ResponseWithError(w, 400, "Bad request")
		return 
	}

	if errors := params.ValidateUser(); len(errors) > 0{
		util.ResponseWithJson(w, 400, errors)
		return 
	}
	
	ctx := context.Background()
	userData, err := s.DB.UpdateUser(ctx, user.ID, params)

	if err != nil {
		s.ErrorLogger.Println(err)
		util.ResponseWithError(w, 400, "User doesn't exist")
		return 
	} 

	util.ResponseWithJson(w, 200, userData)
}