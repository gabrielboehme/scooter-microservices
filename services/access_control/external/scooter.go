package external

import (
	// "os"
	// "io/ioutil"
	"errors"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"

	"access/internal/processors"
)

var scooterServiceUrl string = "http://scooter-app:8000/scooter"

type Scooter struct {
	ID            *string `json:"id"`
	SerialNumber  *string `json:"serial_number"`
	Status        *string `json:"status"`
	State         *string `json:"state"`
	Error         string `json:"error"`
}

func GetScooterOr404(scooterSerialNumber string, w http.ResponseWriter, r *http.Request) *Scooter {
	scooterEP := scooterServiceUrl + "/" + scooterSerialNumber

	resp, err := http.Get(scooterEP)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	// Create a Scooter object from response
	scooter := Scooter{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&scooter); err != nil {
		fmt.Println("Error: ", err.Error())
		processors.RespondError(w, http.StatusBadRequest,
			fmt.Sprintf("Scooter not found with serial number %s", scooterSerialNumber),
		)
		return nil
	}
	
    defer resp.Body.Close()
	
    if resp.StatusCode != http.StatusOK {
		processors.RespondError(w, http.StatusBadRequest,
			fmt.Sprintf("Scooter not found with serial number %s", scooterSerialNumber),
		)
        return nil
    }
	return &scooter
}

func UpdateScooterStatus(scooterSerialNumber string, status string) error {
	scooterEP := scooterServiceUrl + "/" + scooterSerialNumber
	requestBody := []byte(`{"status": "` + status + `"}`)

	req, err := http.NewRequest("PATCH", scooterEP, bytes.NewBuffer(requestBody))
	if err != nil {
        return fmt.Errorf("Internal server error")
    }
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	// Checks if scooter is not in use
	var response struct {
		Error string `json:"error"`
	}
	
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("Error:", err)
		return errors.New("Internal server error")
	}
	if response.Error == "scooter_already_in_use" {
		return errors.New("scooter_already_in_use")
	} 
	if response.Error == "scooter_already_available" {
		return errors.New("scooter_not_rented")
	}
	if response.Error == "scooter_out_of_order" {
		return errors.New("scooter_out_of_order")
	} 

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error in response: %s", response)
		fmt.Printf("Error in response: %d", resp.StatusCode)
		return errors.New("Failed to rent scooter")
	}
	return nil
}