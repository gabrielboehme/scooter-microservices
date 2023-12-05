package server

import (
	"github.com/gorilla/mux"

	"payments/internal/api"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/pay", api.PayForScooter).Methods("POST")
}
