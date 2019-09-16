package slimlist

import (
	"errors"
	"fmt"

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
	t := NewTask(description, u)
	tasks = append(tasks, *t)
	return t, nil
}

//ChangeTaskStatus changes task status
func (u *User) ChangeTaskStatus(task *Task, status int) error {
	if status == Archived && u.Role != AdminStatus {
		return errors.New("Only administrator allowed to archive a task")
	}
	task.ChangeStatus(status)

	return nil
}

//ChangeTaskStatusWithNotify changes task status with email notification
func (u *User) ChangeTaskStatusWithNotify(es EmailSender, task *Task, status int) error {

	err := u.notifyTaskUsers(es, task, u.getNotifyMsg(task, status))
	if err != nil {
		return err
	}

	return u.ChangeTaskStatus(task, status)

}

//AddComment adds Comment to task
func (u *User) AddComment(comment string, t *Task) (*Comment, error) {
	c := NewComment(comment, u)
	t.Comments = append(t.Comments, *c)
	return c, nil
}

//DeleteComment adds Comment to task
func (u *User) DeleteComment(id string, t *Task) error {

	if u.Role != AdminStatus {
		return fmt.Errorf("Only administrator can delete comments")
	}

	for i, c := range t.Comments {
		if c.ID == id {
			t.Comments = append(t.Comments[:i], t.Comments[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Comment with given id:%s couldn't be found", id)
}

func (u *User) notifyTaskUsers(es EmailSender, task *Task, msg string) error {

	body := []byte(msg)
	to := []string{}

	for _, usr := range task.Users {
		if u.Email != usr.Email {
			to = append(to, usr.Email)
		}
	}
	return es.Send(to, body)
}

func (u *User) getNotifyMsg(task *Task, status int) string {
	return fmt.Sprintf("Task ID %s . User %s  Status changed from %s to %s", task.ID, u.Email, StatusTexts[task.Status], StatusTexts[status])
}
