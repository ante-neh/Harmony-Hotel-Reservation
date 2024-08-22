package server

import "net/http"

func (s *Server) Router() http.Handler{
	mux := http.NewServeMux() 
	mux.Handle("GET /api/v1/healthz", http.HandlerFunc(s.HandleHealthz))
	mux.Handle("POST /api/v1/users",http.HandlerFunc(s.HandleCreateUser))
	mux.Handle("GET /api/v1/users", s.authMiddleware(s.HandleGetUser))
	return s.secureHeaders(mux) 
}