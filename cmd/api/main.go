package main

import (
	"context"
	"flag"
	"log"
	"os"
	"github.com/ante-neh/Harmony-Hotel-Reservation/internal/server"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	connection := flag.String("connectionString", connectionString, "database connection string")
	flag.Parse() 


	//create custome loggers
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ltime | log.Ldate)
	errorLogger := log.New(os.Stdout, "Error: ", log.Ltime | log.Ldate | log.Lshortfile) 


	//get mongodb client 
	client, err := openDb(*connection)

	if err != nil{
		errorLogger.Fatal("Unable to Connect to the database", err)
	}


	//create a new server 
	app := server.NewServer(infoLogger, errorLogger, *address, client) 

	//start the server 
	server := app.Start() 
	app.InfoLogger.Println("Server is running on port ", *address) 
	err = server.ListenAndServe()
	
	if err != nil{
		app.ErrorLogger.Println(err)
	}
}


func openDb(connectionString string)(*mongo.Client, error){
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	
	if err != nil{
		return nil, err
	}

	return client, nil  
}