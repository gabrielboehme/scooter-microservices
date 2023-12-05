package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"

	"scooter/internal/processors"
	"scooter/internal/model"
	"strconv"
)


func GetScooter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Inside Get scooter")
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

	latitude := vars["latitude"]
	longitude := vars["logitude"]
	serialNumber := vars["serial_number"]
	scooter := model.GetScooterOr404(serialNumber, w, r)
	if scooter == nil {
		return
	}

	lat, err := strconv.ParseFloat(latitude, 64)
    if err != nil {
        processors.RespondError(w, http.StatusBadRequest, "Invalid latitude")
        return
    }
	lon, err := strconv.ParseFloat(longitude, 64)
    if err != nil {
        processors.RespondError(w, http.StatusBadRequest, "Invalid longitude")
        return
    }

	newLocation := model.Location{
		Latitude: &lat,
		Longitude: &lon,
	}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&newLocation); err != nil {
        processors.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }
    defer r.Body.Close()

	if err := scooter.ValidateLocationChange(&newLocation); err != nil {
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
    vars := mux.Vars(r)
    latitude := vars["latitude"]
    longitude := vars["longitude"]

    // Parse the latitude and longitude from the request into float64 values
    lat, err := strconv.ParseFloat(latitude, 64)
    if err != nil {
        processors.RespondError(w, http.StatusBadRequest, "Invalid latitude")
        return
    }

    lon, err := strconv.ParseFloat(longitude, 64)
    if err != nil {
        processors.RespondError(w, http.StatusBadRequest, "Invalid longitude")
        return
    }

    // Define a radius for the proximity search (adjust this as needed)
    // In this example, we assume a radius of 1 degree
    radius := 0.009

    // Perform a spatial query to find scooters within the specified radius
    var nearScooters []model.Scooter
    result := model.DB.Where(
        "ST_DWithin(location, ST_MakePoint(?, ?)::geography, ?)",
        lon, lat, radius,
    ).Find(&nearScooters)

    if result.Error != nil {
        processors.RespondError(w, http.StatusInternalServerError, result.Error.Error())
        return
    }

    // Serialize the near scooters to a list of Location structs
    locations := make([]model.Location, len(nearScooters))
    for i, scooter := range nearScooters {
        locations[i] = model.SerializeScooterToLocation(&scooter)
    }

    // Respond with the list of near scooters' locations in JSON format
    processors.RespondJSON(w, http.StatusOK, locations)
}