package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bhushan-aruto/db"
	"github.com/bhushan-aruto/repository"
	"github.com/bhushan-aruto/route"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occurred while loading the environment variables, Error -> %v\n", err.Error())
	}

	db, err := db.Connect()

	if err != nil {
		log.Fatalf("error occurred while connecting to the database, Error -> %v\n", err.Error())
	}

	log.Printf("connected to database\n")

	postgresRepository := repository.NewPostgresRepository(db)

	router := route.NewRouter(postgresRepository)

	serverAddress := os.Getenv("SERVER_ADDRESS")

	log.Printf("server is running on: %v\n", serverAddress)
	http.ListenAndServe(serverAddress, router)
}