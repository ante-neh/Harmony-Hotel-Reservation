package main

import (
	"flag"
	"log"
	"os"
	"github.com/ante-neh/Harmony-Hotel-Reservation/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{

	}

	//importing the port address and the connection string from .env file
	port := os.Getenv("PORT")
	connectionString := os.Getenv("CONN")

	//accepting the port address and the connection string from the cli
	address := flag.String("address", port, "Server address")
	_ = flag.String("connectionString", connectionString, "database connection string")
	flag.Parse() 

	//create custome loggers
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ltime | log.Ldate)
	errorLogger := log.New(os.Stdout, "Error: ", log.Ltime | log.Ldate | log.Lshortfile) 


	//create a new server 
	app := server.NewServer(infoLogger, errorLogger, *address) 

	//start the server 
	server := app.Start() 
	app.InfoLogger.Println("Server is running on port ", *address) 
	err = server.ListenAndServe()
	
	if err != nil{
		app.ErrorLogger.Println(err)
	}
}