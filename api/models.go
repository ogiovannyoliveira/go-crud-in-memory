package api

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type ID uuid.UUID

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func (u User) Validate() error {
	if u.FirstName == "" {
		return errors.New("First name is required.")
	}
	if len(u.FirstName) < 2 || len(u.FirstName) > 20 {
		return errors.New("First name should have length between 2 and 20 characters.")
	}

	if u.LastName == "" {
		return errors.New("Last name is required.")
	}
	if len(u.LastName) < 2 || len(u.LastName) > 20 {
		return errors.New("Last name should have length between 2 and 20 characters.")
	}

	if u.Biography == "" {
		return errors.New("Biography is required.")
	}
	if len(u.Biography) < 2 || len(u.Biography) > 20 {
		return errors.New("Biography should have length between 20 and 450 characters.")
	}

	return nil
}

type Application struct {
	Data map[ID]User
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}
