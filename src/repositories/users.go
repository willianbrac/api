package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, email, password) values(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastUserInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastUserInsertID), nil
}