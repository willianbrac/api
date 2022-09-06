package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Conn()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", userID)))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createuser"))
}
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createuser"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createuser"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createuser"))
}