package external

import (
	// "os"
	// "io/ioutil"
	"errors"
	// "strings"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"

	"rents/internal/processors"
)

var paymentsServiceUrl string = "http://payments-app:8000"

type PaymentResponse struct {
	Error        *string `json:"error"`
}

func PayScooterRent(rentId string, cardNumber string, w http.ResponseWriter, r *http.Request) error {
	paymentEP := paymentsServiceUrl + "/pay"
	data := struct {
		CardNumber string `json:"card_number"`
		RentID     string `json:"rent_id"`
	}{
		CardNumber: cardNumber,
		RentID:     rentId,
	}

	// Marshal the data into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return errors.New("Failed to pay - invalid payment info")
	}
	jsonDataReader := bytes.NewReader(jsonData)

	resp, err := http.Post(paymentEP, "application/json", jsonDataReader)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	var response PaymentResponse 
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("\nError:", err)
		return errors.New("Internal server error decoding")
	}
	if response.Error != nil {
		fmt.Printf("\nError: %s", *response.Error)
		processors.RespondError(w, http.StatusBadRequest, "Failed to pay for rent. Try later")
		return errors.New(*response.Error)
	}

	return errors.New("Internal server error")
}