package slimlist

import (
	"testing"

	"github.com/google/uuid"
)

func TestUser(t *testing.T) {

	user := User{
		ID:    uuid.New().String(),
		Email: "test@user.com"}

	if user.Email != "test@user.com" {
		t.Errorf("New task email was incorrect, got: %s, want: test@user.com", user.Email)
	}

}

func TestNewUser(t *testing.T) {
	email := "test@user.com"
	user, err := NewUser(email)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}
	if user.Email != email {
		t.Errorf("New task email was incorrect, got: %s, want: %s", user.Email, email)
	}

}

func TestNewUserEmailDublication(t *testing.T) {
	email := "somedubemail@user.com"
	_, err := NewUser(email)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}
	_, err = NewUser(email)
	if err == nil {
		t.Error("[User] NewUser dublicate email didn't threw any error", err)
	}

}

func TestNewTask(t *testing.T) {

	user := User{
		ID:    uuid.New().String(),
		Email: "test@user.com"}

	desc := "new task description"
	newtask, err := user.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}
	if newtask.Description != desc {
		t.Errorf("Newly created task description was incorrect, got: %s, want: %s", newtask.Description, desc)
	}

	if newtask.Status != New {
		t.Errorf("Newly created task status was incorrect, got: %d, want: %d", newtask.Status, New)
	}

	if newtask.ID == "" {
		t.Errorf("Newly created task ID was incorrect, got: %s", newtask.ID)
	}

	if len(newtask.Users) != 1 {
		t.Errorf("Newly created task users len was incorrect, got: %d, want: %d", len(newtask.Users), 1)
	}

	if newtask.Users[0].Email != user.Email {
		t.Errorf("Newly created task user email was incorrect, got: %s, want: %s", newtask.Users[0].Email, user.Email)
	}

}
