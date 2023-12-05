package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"

	"scooter/users/internal/processors"
	"scooter/users/internal/model"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := []model.User{}
	model.DB.Find(&users)
	processors.RespondJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cpf := vars["cpf"]
	user := model.GetUserOr404(cpf, w, r)
	if user == nil {
		return
	}
	processors.RespondJSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	if err := user.ValidateUser(); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	if err := user.RaiseOnUserExists(); err != nil {
		processors.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DB.Create(&user).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, "internal server error")
		fmt.Println(err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cpf := vars["cpf"]
	user := model.GetUserOr404(cpf, w, r)
	if user == nil {
		return
	}

	// Create a custom JSON decoder for the request body
    decoder := model.NewUserUpdateDecoder(r.Body)

    // Use the custom decoder to update the user struct
	var userUpdated model.UserUpdate
    if err := decoder.Decode(&userUpdated); err != nil {
        processors.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }
	defer r.Body.Close()

	user.UpdateUser(&userUpdated)
	if err := model.DB.Save(&user).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cpf := vars["cpf"]
	user := model.GetUserOr404(cpf, w, r)
	if user == nil {
		return
	}

	if err := model.DB.Delete(&user).Error; err != nil {
		processors.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	processors.RespondJSON(w, http.StatusNoContent, nil)
}
