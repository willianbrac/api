package repositories

import (
	"api/src/models"
	"database/sql"
)

type Books struct {
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *Books {
	return &Books{db}
}

func (repository Books) Create(book models.Book) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO books (title, category, synopsis, author_id) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(book.Title, book.Category, book.Synopsis, book.AuthorID)
	if err != nil {
		return 0, err
	}
	lastUserInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastUserInsertID), nil
}

func (repository Books) List() ([]models.Book, error) {
	rows, err := repository.db.Query("SELECT id, title, category, synopsis, author_id FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Category,
			&book.Synopsis,
			&book.AuthorID,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (repository Books) FindOne(bookID uint64) (models.Book, error) {
	rows, err := repository.db.Query(
		"SELECT title, category, synopsis, author_id FROM books b INNER JOIN users u on u.id = b.author_id WHERE b.id = ?", bookID)
	if err != nil {
		return models.Book{}, err
	}
	defer rows.Close()
	var book models.Book
	if rows.Next(){
		if err = rows.Scan(
			&book.Title,
			&book.Category,
			&book.Synopsis,
			&book.AuthorID,
		); err != nil {
			return models.Book{}, err
		}
	}
	return book, nil
}

func (repository Books) Update(ID uint64, body models.Book) error {
	statement, err := repository.db.Prepare(
		"UPDATE books SET title = ?, category = ?, synopsis = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(body.Title, body.Category, body.Synopsis, ID);err != nil {
		return err
	}
	return nil
}

func (repository Books) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM books WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repository Books) FindByEmail(email string) (models.User, error) {
	rows, err := repository.db.Query("SELECT id, password FROM books WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()
	var user models.User
	if rows.Next(){
		if err = rows.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}