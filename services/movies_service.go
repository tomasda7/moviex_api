package services

import (
	"errors"
	"fmt"

	"github.com/tomasda7/moviex_api/models"
)


func GetMovieDetail(movie_id string) (models.Movie_Detail, error) {

	movieFromApi, err := FetchMovie(movie_id)
	if err != nil {
		return models.Movie_Detail{}, err
	}

	// validate that a movie was found with the ID given
	if movieFromApi.ID == 0 {
		msg := fmt.Sprintf("A movie with the ID %s was not found.", movie_id)
		return models.Movie_Detail{}, errors.New(msg)
	}

	// increment views
	views, err := IncrementViews(movieFromApi.ID, movieFromApi.Title, movieFromApi.Overview)
	if err != nil {
		return models.Movie_Detail{}, err
	}

	// retrieve comments
	comments, err := GetAllComments(movieFromApi.ID)
	if err != nil {
		return models.Movie_Detail{}, err
	}

	// response struct
	movieDetail := models.Movie_Detail{
		Movie_ID: movieFromApi.ID,
		Title: movieFromApi.Title,
		Overview: movieFromApi.Overview,
		Views: views,
		Comments: comments,
	}

	return movieDetail, nil
}

func GetMostViewedMovies() ([]models.Detail, error) {

	rows, err := DB.Query(`SELECT * FROM detail ORDER BY views DESC LIMIT 3;`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := []models.Detail{}

	for rows.Next() {
		movie, err := ScanMovies(rows)

		if err != nil {
			return nil, err
		}
		movies = append(movies, *movie)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return movies, nil
}
