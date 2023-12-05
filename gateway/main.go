package main

import (
	"fmt"
	"net/http"
	"os"
	"net/http/httputil"
	"net/url"
	"github.com/gorilla/mux"
	"strings"
)

// Define a struct to hold information about your microservices
type Service struct {
	Name     string
	ProxyURL string
}

// List of available microservices
var services = []Service{
	{Name: "users", ProxyURL: "http://user-app:8000"},
	{Name: "scooters", ProxyURL: "http://scooter-app:8000"},
	{Name: "rents", ProxyURL: "http://rents-app:8000"},
	// Add more services as needed
}

func main() {
	r := mux.NewRouter()

	// Register a route for each microservice
	for _, service := range services {
		serviceURL, _ := url.Parse(service.ProxyURL)
		proxy := httputil.NewSingleHostReverseProxy(serviceURL)

		// Define a route for each microservice
		r.PathPrefix("/" + service.Name).HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Extract the service name
			parts := strings.SplitN(req.URL.Path, "/", 3)
			fmt.Println("Value of parts[2]:", parts[2])
			if len(parts) < 2 {
				http.Error(w, "Not Found", http.StatusNotFound)
				return
			}

			// Remove the service name from the path
			req.URL.Path = "/" + parts[2]

			// Proxy the request to the corresponding microservice
			proxy.ServeHTTP(w, req)
		})
	}

	// Define a catch-all route for unmatched routes
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	// Create and start the HTTP server
	port := os.Getenv("APP_PORT")
	port = fmt.Sprintf(":%s",port)
	server := &http.Server{
		Addr:    port, // The gateway will listen on port 8080
		Handler: r,
	}

	fmt.Printf("Gateway is listening on %s...", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
