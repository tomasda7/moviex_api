package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tomasda7/moviex_api/services"
)

func GetMovieDetail(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["movie_id"]

	movie, err := services.GetMovieDetail(paramId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error while trying to fetch the movie."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func GetMostViewedMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := services.GetMostViewedMovies()
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error while trying to retrieve the movies from DB."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
