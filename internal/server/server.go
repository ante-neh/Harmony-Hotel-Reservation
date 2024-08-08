package server

import (
	"log"
	"net/http"
)

type Server struct {
	infoLog  *log.Logger
	errorLog *log.Logger 
	address string 
}


func NewServer(infoLog, errorLog *log.Logger, address string) *Server{
	return &Server{
		infoLog:infoLog,
		errorLog:errorLog,
		address:address,
	}
}


func (s *Server) Start() *http.Server{
	return &http.Server{
		Addr:s.address,
		ErrorLog:s.errorLog,
		Handler:s.Router(),
	}
}