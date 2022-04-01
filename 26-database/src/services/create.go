package services

import (
	"database-golang/src/database"
	"database-golang/src/types"

	"errors"
)

func CreateUser(user *types.User) (u *types.User, err error) {
	db, err := database.Connect()
	if err != nil {
		return nil, errors.New("database not connecting")
	}
	defer db.Close()

	statment, err := db.Prepare(`
		INSERT INTO usuarios (first_name, last_name, email, location, is_admin) 
		VALUES (?,?,?,?,?)
	`)
	if err != nil {
		return nil, errors.New("error on creating statement")
	}
	defer statment.Close()

	inserted, err := statment.Exec(user.FirstName, user.LastName, user.Email, "brazil", false)

	if err != nil {
		return nil, errors.New("error on executing statement")
	}

	insertedId, err := inserted.LastInsertId()
	if err != nil {
		return nil, errors.New("error on getting LastInsertId")
	}

	user.ID = int(insertedId)
	return user, nil
}
