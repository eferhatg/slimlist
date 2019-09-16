package slimlist

import "github.com/google/uuid"

//Comment keeps task comments
type Comment struct {
	ID      string
	Comment string
	User    *User
}

//NewComment creates newComment
func NewComment(comment string, u *User) *Comment {
	return &Comment{
		ID:      uuid.New().String(),
		Comment: comment,
		User:    u}
}
