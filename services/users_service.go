package services

import (
	"errors"
	"fmt"

	"github.com/tomasda7/moviex_api/models"
)

func UpdateUserInfo(user_id int, updatedUSer models.User) (string, error) {

 	if updatedUSer.Nickname == "" || updatedUSer.Email == "" || updatedUSer.Password == "" {
		return "", errors.New("if you are not looking to update some fields, enter the current info")
	}

	var userID int

	err := DB.QueryRow("UPDATE users SET nickname=$1, email=$2, password=$3 WHERE id=$4 RETURNING id", updatedUSer.Nickname, updatedUSer.Email, updatedUSer.Password, user_id).Scan(&userID)

	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("The comment user info the ID %d was edited successfully.", userID)
	return msg, nil
}
