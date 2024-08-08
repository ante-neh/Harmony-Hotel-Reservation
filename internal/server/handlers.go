package server

import (
	"net/http"
	"github.com/ante-neh/Harmony-Hotel-Reservation/util"
)

func (s *Server) HandleHealthz(w http.ResponseWriter, r *http.Request){
	util.ResponseWithJson(w, 200, map[string]string{"message":"server is up"})
}