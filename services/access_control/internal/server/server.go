package server

import (
	"github.com/gorilla/mux"

	"access/internal/api"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/lock/{scooter}", api.LockScooter).Methods("POST")
	router.HandleFunc("/unlock/{scooter}", api.UnlockScooter).Methods("POST")
}
