package slimlist

import (
	"testing"

	"github.com/google/uuid"
)

func TestTask(t *testing.T) {

	taskdesc := "New Task"
	email := "TestTask@slimlist.com"

	task := Task{
		ID:          uuid.New().String(),
		Description: taskdesc,
		Status:      New,
		Users: []User{
			User{
				ID:    uuid.New().String(),
				Email: email},
		}}

	if task.Description != taskdesc {
		t.Errorf("New task description was incorrect, got: %s, want: %s", task.Description, taskdesc)
	}

	if task.Status != New {
		t.Errorf("New task status was incorrect, got: %d, want: %d.", task.Status, New)
	}

	if len(task.Users) != 1 {
		t.Errorf("New task users len was incorrect, got: %d, want: %d", len(task.Users), 1)
	}
}

func TestAssignToUser(t *testing.T) {

	taskdesc := "New Task"
	email := "TestAssignToUser@slimlist.com"
	email2 := "TestAssignToUser1@slimlist.com"

	task := Task{
		ID:          uuid.New().String(),
		Description: taskdesc,
		Status:      New,
		Users: []User{
			User{
				ID:    uuid.New().String(),
				Email: email},
		},
	}

	newuser := User{ID: uuid.New().String(), Email: email2}
	err := task.AssignToUser(newuser)

	if err != nil {
		t.Error("[Task] AssignToUser failed", err)
	}

	if len(task.Users) != 2 {
		t.Errorf("User assigned tasks  len was incorrect, got: %d, want: %d", len(task.Users), 2)
	}

	var found bool
	for _, u := range task.Users {
		if u.Email == email2 {
			found = true
		}
	}

	if found != true {
		t.Errorf("Newly assigned user couldn't be found")
	}
}

func TestChangeStatus(t *testing.T) {
	taskdesc := "New Task"
	email := "TestChangeStatus@slimlist.com"

	task := Task{
		ID:          uuid.New().String(),
		Description: taskdesc,
		Status:      New,
		Users: []User{
			User{
				ID:    uuid.New().String(),
				Email: email},
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
