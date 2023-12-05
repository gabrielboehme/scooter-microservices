package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"access/internal/server"
	"access/internal/util"
)

// função principal
func main() {

	if os.Getenv("DEPLOYMENT") != "PROD" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Create the router
	router := mux.NewRouter()

	// Use logging middleware
	router.Use(util.LoggingMiddleware)

	// Register the routes
	server.RegisterRoutes(router)
	port := os.Getenv("APP_PORT")
	fmt.Println(fmt.Sprintf("Started server at: :%s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port ), router))

}
