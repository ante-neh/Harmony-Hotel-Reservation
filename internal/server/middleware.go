package server

import (
	"net/http"

	"github.com/ante-neh/Harmony-Hotel-Reservation/types"
)

type authHandler func (w http.ResponseWriter, r *http.Request, user types.User)

func (s *Server) authMiddleware(next authHandler)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		next(w, r, types.User{})
	}
}


func (s *Server) secureHeaders(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}