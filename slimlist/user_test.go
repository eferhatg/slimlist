package slimlist

import (
	"fmt"
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
		t.Error("[User] NewUser dublicate email didn't threw any error")
	}

}
func TestSetAdmin(t *testing.T) {
	email := "admin@user.com"
	u, err := NewUser(email)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}
	err = u.SetAdmin()
	if err != nil {
		t.Error("[User] SetAdmin failed", err)
	}

	if u.Role != "Admin" {
		t.Errorf("Newly setted admin user role was incorrect, got: %s, want: Admin", u.Role)
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

func TestChangeTaskStatus(t *testing.T) {

	email := "changetaskadmin@user.com"
	desc := "new task description"
	u, err := NewUser(email)

	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	task, err := u.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}

	err = u.ChangeTaskStatus(task, InProgress)
	if err != nil {
		t.Error("[User] ChangeTaskStatus failed", err)
	}

	err = u.ChangeTaskStatus(task, Archived)
	if err == nil {
		t.Error("[User] ChangeTaskStatus with non-admin user didn't threw any error")

	}

	err = u.SetAdmin()
	if err != nil {
		t.Error("[User] SetAdmin failed", err)
	}

	err = u.ChangeTaskStatus(task, Archived)
	if err != nil {
		t.Error("[User] ChangeTaskStatus failed", err)

	}

	if task.Status != Archived {
		t.Error("Newly changed status was not Archived")
	}

}

func TestChangeTaskStatusWithNotify(t *testing.T) {

	f, r := mockSend(nil)
	sender := &emailSender{send: f}

	email := "changetaskwithnotify@user.com"
	desc := "new task description"
	u, err := NewUser(email)

	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	task, err := u.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}

	body := fmt.Sprintf("Task ID: %s \n User:%s \n\n Status changed from %s to %s", task.ID, u.Email, StatusTexts[task.Status], StatusTexts[InProgress])

	err = u.ChangeTaskStatusWithNotify(sender, task, InProgress)
	if err != nil {
		t.Error("[User] ChangeTaskStatus failed", err)
	}

	if string(r.msg) != body {
		t.Errorf("wrong message body.\nexpected: %v\n got: %s", body, r.msg)
	}
	if task.Status != InProgress {
		t.Error("Newly changed status was not InProgress")
	}

}

func TestNotifyTaskUsers(t *testing.T) {
	f, r := mockSend(nil)
	sender := &emailSender{send: f}

	email := "notifytaskusers@user.com"
	email1 := "notifytaskusers1@user.com"

	desc := "new task description"
	body := "msg_body"

	u, err := NewUser(email)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	u1, err := NewUser(email1)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	task, err := u.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}

	err = task.AssignToUser(*u1)
	if err != nil {
		t.Error("[Task] AssignToUser failed", err)
	}

	u.notifyTaskUsers(sender, task, body)
	if err != nil {
		t.Error("[User] notifyTaskUsers failed", err)
	}

	if string(r.msg) != body {
		t.Errorf("wrong message body.\nexpected: %v\n got: %s", body, r.msg)
	}

	to := []string{u1.Email}

	if r.to[0] != to[0] {
		t.Errorf("wrong message body.\nexpected: %v\n got: %s", to[0], r.to[0])
	}

}
