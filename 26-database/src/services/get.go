package services

import (
	"database-golang/src/database"
	"database-golang/src/types"
	"errors"
)

func GetUsuarios() ([]types.User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, errors.New("database not connecting")
	}
	defer db.Close()

	cursor, err := db.Query("SELECT id, first_name, last_name, email FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var users []types.User

	for cursor.Next() {
		var user types.User
		if err := cursor.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, errors.New("error on scanning user")
		}
		users = append(users, user)
	}

	return users, nil
}
