package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Conn()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	
	repository := repositories.NewUsersRepository(db)
	userSavedDatabase, err := repository.FindByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifiPassword(userSavedDatabase.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CreateToken(userSavedDatabase.ID)
	fmt.Println(token)
	w.Write([]byte(token))
}
