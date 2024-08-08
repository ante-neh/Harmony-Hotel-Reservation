package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/ante-neh/Hormony-Hotel-Reservation/api"
	"github.com/joho/godotenv"
)

type application struct{
	errorLog *log.Logger 
	infoLog *log.Logger 
}
func main() {
	godotenv.Load() 
	port := os.Getenv("PORT")
	address := flag.String("address", port, "port address" )
	infoLogger := log.New(os.Stdout, "Info: ", log.Ldate | log.Ltime)
	errorLogger := log.New(os.Stdout, "Error: ", log.Ldate | log.Ltime | log.Lshortfile)

	app := &application{
		errorLog:errorLogger,
		infoLog:infoLogger, 
	}

	server := http.Server{
		Handler:api.rounter,
		Addr:*address,
		ErrorLog:errorLogger,
	}

	err := server.ListenAndServe()
	infoLogger.Println("server is running on port: ", *address)
	if err != nil{
		errorLogger.Fatal("unable to start the server")
	}
	
}