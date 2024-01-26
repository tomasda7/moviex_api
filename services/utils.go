package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tomasda7/moviex_api/models"
)

func FetchMovie(movie_id string) (models.Movie_Api, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=963cd85bee041ada445e94ce84c8e5ff", movie_id)

	res, err := http.Get(url)

	if err != nil {
		return models.Movie_Api{}, err
	}
	defer res.Body.Close()

	var MovieFromApi models.Movie_Api
	json.NewDecoder(res.Body).Decode(&MovieFromApi)

	return MovieFromApi, nil
}

func IncrementViews(movie_id int, title string, overview string) (int, error) {

	if !theMovieExistsInDB(movie_id) {

		stmt, err := DB.Prepare("INSERT INTO detail (movie_id, title, overview, views) VALUES ($1,$2,$3,$4) RETURNING views")

		if err != nil {
			return 0, err
		}
		defer stmt.Close()

		var views int

		err = stmt.QueryRow(movie_id,title,overview,1).Scan(&views)

		if err != nil {
			return 0, err
		}

		return views, nil
	} else {
		var views int

		err := DB.QueryRow("SELECT views FROM detail WHERE movie_id=$1",movie_id).Scan(&views)
		if err != nil {
			return 0, err
		}

		stmt, err := DB.Prepare("UPDATE detail SET views=$1 WHERE movie_id=$2 RETURNING views")
		if err != nil {
			return 0, err
		}
		defer stmt.Close()

		var updViews int

		err = stmt.QueryRow(views+1, movie_id).Scan(&updViews)

		if err != nil {
			return 0, err
		}

		return updViews, nil
	}
}


func theMovieExistsInDB(id int) bool {
	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM detail WHERE movie_id=$1", id).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func GetAllComments(movie_id int) ([]models.Comment_View, error) {

	rows, err := DB.Query(`SELECT users.nickname, comments.content FROM comments INNER JOIN users ON comments.user_id=users.id WHERE comments.movie_id=$1;`,movie_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment_View{}

	for rows.Next() {
		comment, err := scanComments(rows)

		if err != nil {
			return nil, err
		}
		comments = append(comments, *comment)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

type RowScanner interface {
	Scan(dest ...interface{}) error
}

func scanComments(rows RowScanner) (*models.Comment_View, error) {
	var comment models.Comment_View

	err := rows.Scan(&comment.UserName, &comment.Content)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func ScanMovies(rows RowScanner) (*models.Detail, error) {
	var detail models.Detail

	err := rows.Scan(&detail.Movie_ID, &detail.Title, &detail.Overview, &detail.Views)

	if err != nil {
		return nil, err
	}

	return &detail, nil
}




