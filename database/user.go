package database

import (
	"github.com/sarahrajabazdeh/todolist/model"
	"github.com/sarahrajabazdeh/todolist/todolisterrors"
)

func (db *Database) GetUsers() []model.User {
	var users []model.User

	query := "SELECT * FROM public.user"

	// execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		todolisterrors.ThrowError(err)
	}

	defer rows.Close()

	// read the result of the execution
	for rows.Next() {

		user := model.User{}

		// assign to user the result readed
		err := rows.Scan(
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.ID,
		)

		if err != nil {
			todolisterrors.ThrowError(err)
		}

		users = append(users, user)
	}

	return users
}
