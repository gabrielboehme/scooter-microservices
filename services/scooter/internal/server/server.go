package server

import (
	"github.com/gorilla/mux"

	"scooter/internal/api"
)

func RegisterRoutes(router *mux.Router) {
    // Routes with path variables
    router.HandleFunc("/scooter/{serial_number}", api.GetScooter).Methods("GET")
    router.HandleFunc("/scooter/{serial_number}", api.UpdateScooter).Methods("PATCH")
    router.HandleFunc("/scooter/{serial_number}", api.DeleteScooter).Methods("DELETE")

    // Specific routes
    router.HandleFunc("/scooter/{serial_number}/location", api.GetLocation).Methods("GET")
    router.HandleFunc("/scooter/{serial_number}/location", api.UpdateLocation).Methods("PATCH")

    // Routes with path variables should come before more general routes.
    // Be careful with the order and how you define your routes with path variables.
    // Ensure that they don't overlap with other routes.
    
    // Other routes
    router.HandleFunc("/scooter/nearme", api.GetNearScooters).Methods("GET")
    router.HandleFunc("/scooter", api.CreateScooter).Methods("POST")
}
