package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"

	"scooter/internal/processors"
	"scooter/internal/model"
)


func GetScooter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	scooterId := vars["id"]
	scooter := model.GetScooterOr404(scooterId, w, r)
	if scooter == nil {
		return
	}
	processors.RespondJSON(w, http.StatusOK, scooter)
}

func CreateScooter(w http.ResponseWriter, r *http.Request) {
	scooter := model.Scooter{}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&scooter); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	if err := scooter.ValidateScooter(); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DB.Create(&scooter).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, "internal server error")
		fmt.Println(err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusCreated, scooter)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if scooter == nil {
		return
	}
	location := serializers.SerializeScooterToLocation(scooter)

    // Respond with the location data in JSON format
    processors.RespondJSON(w, http.StatusOK, location)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	latitude := vars["latitude"]
	logitude := vars["logitude"]
	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if scooter == nil {
		return
	}

	var newLocation Location
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&newLocation); err != nil {
        processors.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }
    defer r.Body.Close()

	if err := scooter.ValidateLocationChange(newLocation); err != nil {
        processors.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }

	if err := model.DB.Save(&scooter).Error; err != nil {
        processors.RespondError(w, http.StatusInternalServerError, err.Error())
        return
    }
	processors.RespondJSON(w, http.StatusOK, scooter)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if user == nil {
		return
	}

	if err := model.DB.Delete(&scooter).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusNoContent, nil)
}

func GetNearScooters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	latitude := vars["latitude"]
	longitude := vars["longitude"]

	// Implement 
	// Query scooter db to get near me 
	// Return Scooters 

}