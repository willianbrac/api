package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(requestBody, &book); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = json.Unmarshal(requestBody, &book); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	book.AuthorID = userID

	if err = book.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conn()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBooksRepository(db)
	book.ID, err = repository.Create(book)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, book)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	db, err := database.Conn()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	respository := repositories.NewBooksRepository(db)
	books, err := respository.List()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["bookId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	
	db, err := database.Conn()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	respository := repositories.NewBooksRepository(db)
	book, err := respository.FindOne(bookID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["bookId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conn()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBooksRepository(db)
	bookSavedOnDatabase, err := repository.FindOne(bookID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if bookSavedOnDatabase.AuthorID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("Não é possível atualizar um livro de outro usuário"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(body, &book); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = book.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(bookID, book); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["bookId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conn()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewBooksRepository(db)
	bookSavedOnDatabase, err := repository.FindOne(bookID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if bookSavedOnDatabase.AuthorID != userID {
		responses.Err(w, http.StatusForbidden, errors.New("Não é possível deletar uma livro de outro usuário"))
		return
	}

	respository := repositories.NewBooksRepository(db)
	if err = respository.Delete(bookID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}