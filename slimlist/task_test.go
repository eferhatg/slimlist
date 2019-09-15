package slimlist

import (
	"testing"

	"github.com/google/uuid"
)

func TestTask(t *testing.T) {

	task := Task{
		ID:          uuid.New().String(),
		Description: "New Task",
		Status:      New,
		Users: []User{
			User{
				ID:    uuid.New().String(),
				Email: "test@user.com"},
		}}

	if task.Description != "New Task" {
		t.Errorf("New task description was incorrect, got: %s, want: New Task.", task.Description)
	}

	if task.Status != New {
		t.Errorf("New task status was incorrect, got: %d, want: %d.", task.Status, New)
	}

	if len(task.Users) != 1 {
		t.Errorf("New task users len was incorrect, got: %d, want: %d", len(task.Users), 1)
	}
}

func TestAssignToUser(t *testing.T) {
	task := Task{
		ID:          uuid.New().String(),
		Description: "New Task",
		Status:      New,
		Users: []User{
			User{
				ID:    uuid.New().String(),
				Email: "test@user.com"},
		},
	}

	newuseremail := "newuser@user.com"
	newuser := User{ID: uuid.New().String(), Email: newuseremail}
	err := task.AssignToUser(newuser)

	if err != nil {
		t.Error("[Task] AssignToUser failed", err)
	}

	if len(task.Users) != 2 {
		t.Errorf("User assigned tasks  len was incorrect, got: %d, want: %d", len(task.Users), 2)
	}

	var found bool
	for _, u := range task.Users {
		if u.Email == newuseremail {
			found = true
		}
	}

	if found != true {
		t.Errorf("Newly assigned user couldn't be found")
	}
}

func TestChangeStatus(t *testing.T) {
	task := Task{
		ID:          uuid.New().String(),
		Description: "New Task",
		Status:      New,
		Users: []User{
			User{
				ID:    uuid.New().String(),
				Email: "test@user.com"},
		},
	}

	err := task.ChangeStatus(InProgress)
	if err != nil {
		t.Error("[Task] ChangeStatus failed", err)
	}
	if task.Status != InProgress {
		t.Errorf("Task status was incorrect, got: %d, want: %d", task.Status, InProgress)
	}

	err = task.ChangeStatus(Completed)
	if err != nil {
		t.Error("[Task] ChangeStatus failed", err)
	}
	if task.Status != Completed {
		t.Errorf("Task status was incorrect, got: %d, want: %d", task.Status, Completed)
	}

}
