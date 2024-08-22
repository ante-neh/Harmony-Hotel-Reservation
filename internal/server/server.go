package server

import (
	"log"
	"net/http"

	"github.com/ante-neh/Harmony-Hotel-Reservation/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger 
	Address string 
	DB *database.MongoDb
}


func NewServer(infoLog, errorLog *log.Logger, address string, client *mongo.Client) *Server{
	return &Server{
		InfoLogger:infoLog,
		ErrorLogger:errorLog,
		Address:address,
		DB:&database.MongoDb{
			Client:client,
		},
	}
}


func (s *Server) Start() *http.Server{
	return &http.Server{
		Addr:s.Address,
		ErrorLog:s.ErrorLogger,
		Handler:s.Router(),
	}
}