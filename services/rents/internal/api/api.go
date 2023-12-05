package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"

	"rents/internal/processors"
	"rents/internal/model"
	"rents/external"
)


func RentScooter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scooterSerialNumber := vars["scooter"]

	scooter := external.GetScooterOr404(scooterSerialNumber, w, r)
	if scooter == nil {
		processors.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if err := external.UpdateScooterStatus(scooterSerialNumber, "IN_USE"); err != nil {
		processors.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	currentTime := time.Now()
	rent := model.Rent{
		ScooterSerialNumber: &scooterSerialNumber,
		RentStart: &currentTime,
	}
	if err := model.DB.Create(&rent).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	processors.RespondJSON(w, http.StatusOK, rent)
}

func FinishScooterRent(w http.ResponseWriter, r *http.Request) {}