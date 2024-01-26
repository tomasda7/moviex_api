package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tomasda7/moviex_api/models"
	"github.com/tomasda7/moviex_api/services"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["movie_id"]

	id, err := strconv.Atoi(paramId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newComment models.Comment

	err = json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := services.AddComment(1, id, newComment.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to add a new comment."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func RemoveComment(w http.ResponseWriter, r *http.Request) {
	paramUserId := mux.Vars(r)["user_id"]

	user_id, err := strconv.Atoi(paramUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paramCommentId := mux.Vars(r)["comment_id"]

	comment_id, err := strconv.Atoi(paramCommentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := services.RemoveComment(comment_id, user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to delete a comment."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func EditComment(w http.ResponseWriter, r *http.Request) {
	paramUserId := mux.Vars(r)["user_id"]

	user_id, err := strconv.Atoi(paramUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paramCommentId := mux.Vars(r)["comment_id"]

	comment_id, err := strconv.Atoi(paramCommentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newComment models.Comment

	err = json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := services.EditComment(comment_id, user_id, newComment.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("An error occurred while trying to edit a comment."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
