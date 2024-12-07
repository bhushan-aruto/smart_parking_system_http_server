package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bhushan-aruto/cache"
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

	cache, err := cache.Connect()

	if err != nil {
		log.Fatalf("failed to connect to redis, Error -> %v\n", err)
	}

	log.Printf("conneced to database")

	postgresRepository := repository.NewPostgresRepository(db)

	redisRepository := repository.NewRedisRepository(cache)

	router := route.NewRouter(postgresRepository, redisRepository)

	serverAddress := os.Getenv("SERVER_ADDRESS")

	log.Printf("server is running on: %v\n", serverAddress)
	http.ListenAndServe(serverAddress, router)
}
