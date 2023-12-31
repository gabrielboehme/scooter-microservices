package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"

	"access/internal/processors"
	"access/external"
)


func LockScooter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scooterSerialNumber := vars["scooter"]

	scooter := external.GetScooterOr404(scooterSerialNumber, w, r)
	if scooter == nil {
		processors.RespondError(w, http.StatusInternalServerError, fmt.Sprint("Couldnt find scooter %s", scooterSerialNumber))
		return
	}

	if *scooter.Status == "IN_USE" {
		processors.RespondSuccess(w, http.StatusOK, "Scooter locked")
		fmt.Println("\nScooter LOCKED")
		return
	}
	processors.RespondError(w, http.StatusBadRequest, "\nScooter must be in_use to lock it.")
}

func UnlockScooter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scooterSerialNumber := vars["scooter"]

	scooter := external.GetScooterOr404(scooterSerialNumber, w, r)
	if scooter == nil {
		processors.RespondError(w, http.StatusInternalServerError, fmt.Sprint("Couldnt find scooter %s", scooterSerialNumber))
		return
	}

	if *scooter.Status == "AVAILABLE" {
		processors.RespondSuccess(w, http.StatusOK, "Scooter unlocked")
		fmt.Println("\nScooter UNLOCKED")
		return
	}
	processors.RespondError(w, http.StatusBadRequest, "\nScooter must be available to unlock it.")
}