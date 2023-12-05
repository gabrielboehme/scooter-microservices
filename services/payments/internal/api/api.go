package api

import (
	"net/http"
	// "github.com/gorilla/mux"
	"fmt"
	"encoding/json"

	"payments/internal/processors"
	"payments/internal/model"
)


func PayForScooter(w http.ResponseWriter, r *http.Request) {
	payment := model.Payment{}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		processors.RespondError(w, http.StatusBadRequest, "Internal server error")
		return
	}
	defer r.Body.Close()
	
	// Missing validation to check if rent_id exists
	if err := payment.ValidatePayment(); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DB.Create(&payment).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, "internal server error")
		fmt.Println(err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusCreated, payment)
}