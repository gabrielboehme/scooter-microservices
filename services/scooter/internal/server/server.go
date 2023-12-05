package server

import (
	"github.com/gorilla/mux"

	"scooter/internal/api"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/scooter/nearme", api.GetNearScooters).Methods("GET")
	router.HandleFunc("/scooter", api.CreateScooter).Methods("POST")
	router.HandleFunc("/scooter/{serial_number}", api.GetScooter).Methods("GET")
	router.HandleFunc("/scooter/{serial_number}", api.GetScooter).Methods("GET")
	router.HandleFunc("/scooter/{serial_number}", api.DeleteScooter).Methods("DELETE")
	router.HandleFunc("/scooter/{serial_number}/location", api.GetLocation).Methods("GET")
	router.HandleFunc("/scooter/{serial_number}/location", api.UpdateLocation).Methods("PATCH")
}