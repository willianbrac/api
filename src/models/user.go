package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64 `json: "id,omitempty"`
	Name      string `json: "name,omitempty"`
	Email     string `json: "email,omitempty"`
	Password  string `json: "password,omitempty"`
	CreatedAt time.Time `json: "createdAt,omitempty"`
}

func (user *User) Prepare() error{
	if err := user.validate(); err != nil {
		return err
	}

	user.formate()
	return nil
}

func (user *User) validate()  error{
	if user.Name == "" {
		return errors.New("O Name é obrigatório e deve ser informado!")
	}
	if user.Email == "" {
		return errors.New("O Email é obrigatório e deve ser informado!")
	}
	if user.Password == "" {
		return errors.New("A senha é obrigatória e deve ser informada!")
	}
	return nil
}



func (user *User) formate(){
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}