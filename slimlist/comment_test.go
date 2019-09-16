package slimlist

import (
	"testing"
)

func TestNewComment(t *testing.T) {
	email := "TestNewComment@slimlist.com"
	commText := "TestNewComment Comment"

	user, err := NewUser(email)
	if err != nil {
		t.Error("[User] NewUser failed", err)
	}

	c := NewComment(commText, user)
	if c.User.Email != email {
		t.Errorf("New comment user was incorrect, got: %s, want: %s", c.User.Email, email)
	}

	if c.Comment != commText {
		t.Errorf("New comment text was incorrect, got: %s, want: %s", c.Comment, commText)
	}

}
