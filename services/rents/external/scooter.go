package external

import (
	// "os"
	"io/ioutil"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"

	"rents/internal/processors"
)

var scooterServiceUrl string = "http://scooter-app:8000/scooter"

type Scooter struct {
	ID            *string `json:"id"`
	SerialNumber  *string `json:"serial_number"`
	Status        *string `json:"status"`
	State         *string `json:"state"`
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
        return fmt.Errorf("Scooter not found with serial number %s", scooterSerialNumber)
    }
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	fmt.Println("Status Code:", resp.Status)

	// Read and print the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle the error
		fmt.Println("Error reading response body:", err)
		return nil
	}
	fmt.Println("Response Body:", string(responseBody))
	
	if err != nil {
		return fmt.Errorf("Internal server error")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Scooter not found with serial number %s", scooterSerialNumber)
	}
	return nil
}