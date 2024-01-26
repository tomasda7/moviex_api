package services

import (
	"errors"
	"fmt"
)

func AddComment(user_id int, movie_id int, content string) (string, error) {

	stmt, err := DB.Prepare("INSERT INTO comments (user_id, movie_id, content) VALUES ($1,$2,$3) RETURNING movie_id")

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	var movieID int

	err = stmt.QueryRow(user_id, movie_id, content).Scan(&movieID)
	if err != nil {
		return "", err
	}

	var title string

	err = DB.QueryRow("SELECT title FROM detail WHERE movie_id=$1", movieID).Scan(&title)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("The comment was added successfully to the movie '%s'.", title)
	return msg, nil
}

func RemoveComment(comment_id int, user_id int) (string, error) {
	var movie_id int

	err := DB.QueryRow("DELETE FROM comments WHERE comment_id=$1 AND user_id=$2 RETURNING movie_id", comment_id, user_id).Scan(&movie_id)

	if err != nil {
		return "", err
	}

	var title string

	err = DB.QueryRow("SELECT title FROM detail WHERE movie_id=$1", movie_id).Scan(&title)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("The comment was removed successfully from the movie '%s'.", title)
	return msg, nil
}

func EditComment(comment_id int, user_id int, newContent string) (string, error) {

	if newContent == "" {
		return "", errors.New("the comment shouldn't be empty")
	}

	var CommentID int

	err := DB.QueryRow("UPDATE comments SET content=$1 WHERE comment_id=$2 AND user_id=$3 RETURNING comment_id", newContent, comment_id, user_id).Scan(&CommentID)

	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("The comment with the ID %d was edited successfully.", CommentID)
	return msg, nil
}
