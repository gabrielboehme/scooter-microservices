package external

import (
	// "os"
	// "io/ioutil"
	"errors"
	"strings"
	"fmt"
	"net/http"
	// "bytes"
	"encoding/json"

	"rents/internal/processors"
)

var accessControlServiceUrl string = "http://access_control-app:8000"

type AccessControlResponse struct {
	Error        *string `json:"error"`
	Success       *string `json:"success"`
}

func LockScooterOr404(scooterSerialNumber string, w http.ResponseWriter, r *http.Request) error {
	scooterEP := accessControlServiceUrl + "/lock/" + scooterSerialNumber

	resp, err := http.Post(scooterEP, "application/json", strings.NewReader(""))
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	var response AccessControlResponse 
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("\nError:", err)
		return errors.New("Internal server error decoding")
	}
	if response.Error != nil {
		fmt.Printf("\nError: %s", *response.Error)
		processors.RespondError(w, http.StatusBadRequest, "Failed to lock scooter. Try later")
		return errors.New(*response.Error)
	}
	
	if response.Success != nil {
		return nil
	}
	return errors.New("Internal server error")

}

func UnlockScooterOr404(scooterSerialNumber string, w http.ResponseWriter, r *http.Request) error {
	scooterEP := accessControlServiceUrl + "/unlock/" + scooterSerialNumber

	resp, err := http.Post(scooterEP, "application/json", strings.NewReader(""))
	if err != nil {
		fmt.Println("\nError: ", err)
		return err
	}

	var response AccessControlResponse 
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("\nError:", err)
		return errors.New("Internal server error decoding")
	}

	if response.Error != nil {
		fmt.Printf("Error: %s", *response.Error)
		processors.RespondError(w, http.StatusBadRequest, "Failed to unlock scooter. Try later")
		return errors.New(*response.Error)
	}

	if response.Success != nil {
		return nil
	}
	return errors.New("Internal server error")

}