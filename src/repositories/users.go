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

func (repository Users) List() ([]models.User, error) {
	rows, err := repository.db.Query("SELECT id, name, email, createdAt FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository Users) FindOne(userId uint64) (models.User, error) {
	rows, err := repository.db.Query("SELECT id, name, email, createdAt FROM users WHERE id = ?", userId)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()
	var user models.User
	if rows.Next(){
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}