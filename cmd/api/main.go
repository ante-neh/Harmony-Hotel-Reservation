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
		log.Panic()
	}
	port := os.Getenv("PORT")
	connString := os.Getenv("CONN")
	address := flag.String("address", port, "server address")
	dns := flag.String("connectionString", connString, "connection string to mongoDB")
	flag.Parse() 

	infoLogger := log.New(os.Stdout, "INFO: ", log.Ltime | log.Ldate)
	errorLogger := log.New(os.Stdout, "ERROR: ", log.Ltime | log.Ldate | log.Lshortfile) 

	app := server.NewServer(infoLogger, errorLogger, *address )
	server := app.Start() 
	err = server.ListenAndServe()

	log.Println("Server is running on port:", *address)
	
	if err != nil{
		errorLogger.Fatal("Unable to start")
	} 

}