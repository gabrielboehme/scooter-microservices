package api

import (
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request)    {}
func CreateUser(w http.ResponseWriter, r *http.Request) {}
func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// db := r.Context().Value("db").(*sql.DB)

	w.WriteHeader(http.StatusOK)
	// w.Write("Testing")
}
