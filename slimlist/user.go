package slimlist

import (
	"errors"

	"github.com/google/uuid"
)

var users []User

//Defines user roles
const (
	AdminStatus = "Admin"
	UserStatus  = "User"
)

//User keeps user definition
type User struct {
	ID    string
	Email string
	Role  string
}

//NewUser creates new user
func NewUser(email string) (*User, error) {
	newuser := &User{ID: uuid.New().String(), Email: email}

	for _, u := range users {
		if u.Email == email {
			return nil, errors.New("dublicate email error")
		}
	}

	users = append(users, *newuser)
	return newuser, nil
}

//SetAdmin makes user admin
func (u *User) SetAdmin() error {

	if u.Role != AdminStatus {
		u.Role = AdminStatus
	}
	return nil
}

//NewTask creates new task
func (u *User) NewTask(description string) (*Task, error) {
	t := &Task{ID: uuid.New().String(), Description: description, Status: New, Users: []User{*u}}
	tasks = append(tasks, *t)
	return t, nil
}

//ChangeTaskStatus changes task status
func (u *User) ChangeTaskStatus(task *Task, status int) error {
	if status == Archived && u.Role != AdminStatus {
		return errors.New("Only administrator allowed to archive a task")
	}
	task.ChangeStatus(Archived)
	return nil
}
