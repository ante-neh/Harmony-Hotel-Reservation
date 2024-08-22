package server

import (
	"log"
	"net/http"
)

type Server struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger 
	Address string 
}


func NewServer(infoLog, errorLog *log.Logger, address string) *Server{
	return &Server{
		InfoLogger:infoLog,
		ErrorLogger:errorLog,
		Address:address,
	}
}


func (s *Server) Start() *http.Server{
	return &http.Server{
		Addr:s.Address,
		ErrorLog:s.ErrorLogger,
		Handler:s.Router(),
	}
}