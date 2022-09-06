package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64 `json: "id,omitempty"`
	Name      string `json: "name,omitempty"`
	Email     string `json: "email,omitempty"`
	Password  string `json: "password,omitempty"`
	CreatedAt time.Time `json: "createdAt,omitempty"`
}

func (user *User) Prepare(stage string) error{
	if err := user.validate(stage); err != nil {
		return err
	}
	if err := user.formate(stage); err != nil {
		return err
	}
	return nil
}
func (user *User) validate(stage string)  error{
	if user.Name == "" {
		return errors.New("O Name é obrigatório e deve ser informado!")
	}
	if user.Email == "" {
		return errors.New("O Email é obrigatório e deve ser informado!")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Verifique o endereço do seu email!")
	}

	if stage == "post" && user.Password == "" {
		return errors.New("A senha é obrigatória e deve ser informada!")
	}
	return nil
}
func (user *User) formate(stage string) error{
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if stage == "post" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}

	return nil
}