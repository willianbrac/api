package models

import (
	"errors"
	"strings"
	"time"
)

type Book struct {
	ID        	uint64    `json: "id,omitempty"`
	Title      	string    `json: "title,omitempty"`
	Category    string    `json: "category,omitempty"`
	Synopsis  	string    `json: "synopsis,omitempty"`
	AuthorID  	uint64    `json: "authorId,omitempty"`
	CreatedAt 	time.Time `json: "createdAt,omitempty"`
}

func (book *Book) Prepare() error{
	if err := book.validate(); err != nil {
		return err
	}
	if err := book.formate(); err != nil {
		return err
	}
	return nil
}

func (book *Book) validate()  error{
	if book.Title == "" {
		return errors.New("O título é obrigatório e deve ser informado!")
	}
	if book.Category == "" {
		return errors.New("A Categoria é obrigatória e deve ser informada!")
	}
	if book.Synopsis == "" {
		return errors.New("A Sinopse é obrigatória e deve ser informada!")
	}
	return nil
}

func (book *Book) formate() error{
	book.Title = strings.TrimSpace(book.Title)
	book.Category = strings.TrimSpace(book.Category)
	book.Synopsis = strings.TrimSpace(book.Synopsis)
	return nil
}