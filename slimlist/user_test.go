package slimlist

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestUser(t *testing.T) {

	email := "TestUser@slimlist.com"
	user := User{
		ID:    uuid.New().String(),
		Email: email}

	if user.Email != email {
		t.Errorf("New task user was incorrect, got: %s, want: %s", user.Email, email)
	}

}

func TestNewUser(t *testing.T) {
	email := "TestNewUser@slimlist.com"
	user, err := NewUser(email)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}
	if user.Email != email {
		t.Errorf("New task user was incorrect, got: %s, want: %s", user.Email, email)
	}

}

func TestNewUserEmailDublication(t *testing.T) {
	email := "TestNewUserEmailDublication@slimlist.com"
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
	email := "TestSetAdmin@slimlist.com"
	u, err := NewUser(email)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}
	err = u.SetAdmin()
	if err != nil {
		t.Error("[User] SetAdmin failed", err)
	}

	if u.Role != AdminStatus {
		t.Errorf("Newly setted admin user role was incorrect, got: %s, want: %s", u.Role, AdminStatus)
	}

}
func TestNewTask(t *testing.T) {
	email := "TestNewTask@slimlist.com"
	desc := "task desc"
	user := User{
		ID:    uuid.New().String(),
		Email: email}

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

	if newtask.Users[0].Email != email {
		t.Errorf("Newly created task user email was incorrect, got: %s, want: %s", newtask.Users[0].Email, email)
	}

}

func TestChangeTaskStatus(t *testing.T) {

	email := "TestChangeTaskStatus@slimlist.com"
	desc := "task description"

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

	email := "TestChangeTaskStatusWithNotify@slimlist.com"
	desc := "task description"
	u, err := NewUser(email)

	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	task, err := u.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}

	body := u.getNotifyMsg(task, InProgress)
	err = u.ChangeTaskStatusWithNotify(sender, task, InProgress)
	if err != nil {
		t.Error("[User] ChangeTaskStatus failed", err)
	}

	if string(r.msg) != body {
		t.Errorf("wrong message body.\nexpected: %v\n got: %s", body, r.msg)
	}
	if task.Status != InProgress {
		t.Errorf("Newly changed status was not %d", InProgress)
	}

}

func TestNotifyTaskUsers(t *testing.T) {
	f, r := mockSend(nil)
	sender := &emailSender{send: f}

	email := "TestNotifyTaskUsers@slimlist.com"
	email1 := "TestNotifyTaskUsers1@slimlist.com"

	desc := "task description"
	body := "msg body"

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

	if r.to[0] != u1.Email {
		t.Errorf("wrong message body.\nexpected: %v\n got: %s", u1.Email, r.to[0])
	}

}

func TestGetNotifyMsg(t *testing.T) {
	email := "TestGetNotifyMsg@slimlist.com"
	desc := "task description"
	u, err := NewUser(email)

	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	task, err := u.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}

	msg := u.getNotifyMsg(task, InProgress)

	if !strings.Contains(msg, email) ||
		!strings.Contains(msg, task.ID) ||
		!strings.Contains(msg, StatusTexts[InProgress]) {
		t.Error("[User] getNotifyMsg generated wrong msg. got:", msg)

	}
}

func TestAddComment(t *testing.T) {
	email := "TestAddComment@slimlist.com"
	desc := "task description"
	commText := "comment text"
	u, err := NewUser(email)

	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	task, err := u.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}

	c, err := u.AddComment(commText, task)
	if err != nil {
		t.Error("[User] Addcomment failed", err)
	}

	if c.Comment != commText {
		t.Errorf("wrong comment text.\nexpected: %v\n got: %s", commText, c.Comment)

	}

	if len(task.Comments) == 0 {
		t.Error("Newly added comment didn't append to task")
	}
	if task.Comments[0].Comment != commText {
		t.Errorf("wrong comment text.\nexpected: %v\n got: %s", commText, task.Comments[0].Comment)
	}
}

func TestDeleteComment(t *testing.T) {
	email := "TestDeleteComment@slimlist.com"
	desc := "new task description"
	u, err := NewUser(email)

	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	task, err := u.NewTask(desc)
	if err != nil {
		t.Error("[User] Newtask failed", err)
	}

	c, err := u.AddComment("comment", task)
	if err != nil {
		t.Error("[User] AddComment failed", err)
	}

	err = u.DeleteComment(c.ID, task)
	if err == nil {
		t.Error("Regular user shouldn't delete comment")
	}

	u.SetAdmin()

	err = u.DeleteComment("different-id", task)
	if err == nil {
		t.Error("Should throw error after different-id")
	}

	err = u.DeleteComment(c.ID, task)
	if err != nil {
		t.Error("Administrator should delete comment")
	}

}
