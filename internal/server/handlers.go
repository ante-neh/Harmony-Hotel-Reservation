package server

import (
	"encoding/json"
	"net/http"

	"github.com/ante-neh/Harmony-Hotel-Reservation/types"
	"github.com/ante-neh/Harmony-Hotel-Reservation/util"
)

func (s *Server) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	util.ResponseWithJson(w, 200, map[string]string{"message": "server is up"})
}


func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		UserName string `bson:"name" json:"name"`
		Email    string `bson:"email" json:"email"`
		Password string `bson:"password" json:"password"`
	}

	params := &parameter{} 

	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil{
		util.ResponseWithError(w, 400, "Bad request")
		return 
	}

	userData, err := s.DB.CreateUser(params.UserName, params.Email, params.Password)

	if err != nil{
		util.ResponseWithError(w, 400, "Couldn't create user")
	}

	util.ResponseWithJson(w, 201, userData)
}


func (s *Server) HandleGetUser(w http.ResponseWriter, r *http.Request, user types.User){
	
}