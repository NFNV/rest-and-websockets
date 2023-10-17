package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"rest-ws/handlers"
	"rest-ws/server"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
	}

	// Retrieve configuration values from environment variables
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// Create a new server instance using configuration settings
	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Start the server and bind routes to the router
	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	// Define routes and associate them with appropriate handlers
	// In this case, a route for the root path ("/") is defined, and it's associated with the HomeHandler.
	// The route is set to respond to HTTP GET requests.
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
