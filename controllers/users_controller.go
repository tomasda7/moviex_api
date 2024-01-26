package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tomasda7/moviex_api/models"
	"github.com/tomasda7/moviex_api/services"
)

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	paramUserId := mux.Vars(r)["user_id"]

	user_id, err := strconv.Atoi(paramUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedUSer models.User

	err = json.NewDecoder(r.Body).Decode(&updatedUSer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := services.UpdateUserInfo(user_id, updatedUSer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to update user info."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
