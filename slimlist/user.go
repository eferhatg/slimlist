package slimlist

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var users []User

//User keeps user definition
type User struct {
	ID    string
	Email string
}

//NewUser creates new user
func NewUser(email string) (*User, error) {
	newuser := &User{ID: uuid.New().String(), Email: email}
	fmt.Println(users)
	for _, u := range users {
		if u.Email == email {
			return nil, errors.New("dublicate email error")
		}
	}

	users = append(users, *newuser)
	return newuser, nil
}

//NewTask creates new task
func (u *User) NewTask(description string) (*Task, error) {
	return &Task{ID: uuid.New().String(), Description: description, Status: New, Users: []User{*u}}, nil
}
