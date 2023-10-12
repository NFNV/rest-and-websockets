package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Config struct holds the configuration parameters for the server.
type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

// Server is an interface that defines a method to retrieve the server's configuration.
type Server interface {
	Config() *Config
}

// Broker struct represents the server instance.
type Broker struct {
	config *Config
	router *mux.Router
}

// Config returns the configuration of the server.
func (b *Broker) Config() *Config {
	return b.config
}

// NewServer creates a new server instance based on the provided configuration.
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("database is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

// Start initiates the server by binding routes to the router and listening for incoming requests.
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	// Create a new router for this server instance.
	b.router = mux.NewRouter()
	// Bind routes to the router.
	binder(b, b.router)
	log.Println("Starting server on", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
