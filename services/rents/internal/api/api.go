package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"

	"rents/internal/processors"
	"rents/internal/model"
	"rents/external"
)


func StartScooterRent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scooterSerialNumber := vars["scooter"]

	scooter := external.GetScooterOr404(scooterSerialNumber, w, r)
	if scooter == nil {
		processors.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if err := external.UnlockScooterOr404(scooterSerialNumber, w, r); err != nil {
		return
	}
	if err := external.UpdateScooterStatus(scooterSerialNumber, "IN_USE"); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	currentTime := time.Now()
	rent := model.Rent{
		ScooterSerialNumber: &scooterSerialNumber,
		RentStart: &currentTime,
	}


	if err := model.DB.Create(&rent).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	processors.RespondJSON(w, http.StatusOK, rent)
}

func FinishScooterRent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scooterSerialNumber := vars["scooter"]
	rent := model.GetInUseRentOr404(scooterSerialNumber, w, r)
	
	if err := external.LockScooterOr404(scooterSerialNumber, w, r); err != nil {
		return
	}

	if err := external.UpdateScooterStatus(scooterSerialNumber, "AVAILABLE"); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Hardcoded card number
	if err := external.PayScooterRent(rent.ID.String(), "123456789", w, r); err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	

	currentTime := time.Now()
	rent.RentFinish = &currentTime
	if err := model.DB.Save(&rent).Error; err != nil {
		fmt.Printf("Error: %s", err.Error())
		processors.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	processors.RespondJSON(w, http.StatusOK, rent)
}