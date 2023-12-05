package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"strconv"

	"scooter/internal/processors"
	"scooter/internal/model"
	"scooter/internal/util"
)


func GetScooter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
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

func UpdateScooter(w http.ResponseWriter, r *http.Request) {
	// Updates Scooter Fields:
	// Status
	// State 
	vars := mux.Vars(r)
	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if scooter == nil {
		return
	}

	var scooterUpdated model.ScooterUpdate
	decoder := model.NewScooterUpdateDecoder(r.Body)
    if err := decoder.Decode(&scooterUpdated); err != nil {
        processors.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }
	defer r.Body.Close()

	// Tries to update scooter
	if err := scooter.UpdateScooter(&scooterUpdated); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DB.Save(&scooter).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusOK, scooter)


}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if scooter == nil {
		return
	}
	location := model.SerializeScooterToLocation(scooter)

    // Respond with the location data in JSON format
    processors.RespondJSON(w, http.StatusOK, location)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if scooter == nil {
		return
	}

	var locationUpdated model.LocationUpdate
    decoder := model.NewLocationUpdateDecoder(r.Body)
    if err := decoder.Decode(&locationUpdated); err != nil {
        processors.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }
    defer r.Body.Close()

	// Tries to update scooter
	if err := scooter.ValidateLocationChange(&locationUpdated); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DB.Save(&scooter).Error; err != nil {
        processors.RespondError(w, http.StatusInternalServerError, err.Error())
        return
    }
	processors.RespondJSON(w, http.StatusOK, scooter)
}

func DeleteScooter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if scooter == nil {
		return
	}

	if err := model.DB.Delete(&scooter).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusNoContent, nil)
}

func GetNearScooters(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.Query()
    latitudeStr := queryParams.Get("latitude")
    longitudeStr := queryParams.Get("longitude")

    lat, err := strconv.ParseFloat(latitudeStr, 64)
    if err != nil {
        processors.RespondError(w, http.StatusBadRequest, "Invalid latitude")
        return
    }

    lon, err := strconv.ParseFloat(longitudeStr, 64)
    if err != nil {
        processors.RespondError(w, http.StatusBadRequest, "Invalid longitude")
        return
    }

    // Define the radius in kilometers (e.g., 1 kilometer)
    radiusKm := 5.0

    // Calculate the latitude and longitude range using the CalculateRange function
    latitudeMin, latitudeMax, longitudeMin, longitudeMax := util.CalculateLocationMinMaxRange(lat, lon, radiusKm)

    var nearScooters []model.Scooter

    // Execute the GORM query to find available scooters within the specified range
    result := model.DB.Where(
		`latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?
		AND status = 'AVAILABLE'
		AND state = 'ON'`,
		latitudeMin, latitudeMax, longitudeMin, longitudeMax,
	).Find(&nearScooters)

    if result.Error != nil {
        processors.RespondError(w, http.StatusInternalServerError, result.Error.Error())
        return
    }

    // Respond with the list of near available scooters in JSON format
    processors.RespondJSON(w, http.StatusOK, nearScooters)
}
