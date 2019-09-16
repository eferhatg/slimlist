package slimlist

import "github.com/google/uuid"

//Keeps task status
const (
	New = iota
	InProgress
	Completed
	Archived
)

//StatusTexts keeps status explanations
var StatusTexts = []string{
	New:        "New",
	InProgress: "In Progress",
	Completed:  "Completed",
	Archived:   "Archived",
}

var tasks []Task

//Task keeps task definition
type Task struct {
	ID          string
	Description string
	Status      int
	Users       []User
	Comments    []Comment
}

//NewTask creates new task
func NewTask(desc string, u *User) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Description: desc,
		Status:      New,
		Users:       []User{*u},
	}
}

//AssignToUser assigns task to user
func (t *Task) AssignToUser(u User) error {
	t.Users = append(t.Users, u)
	return nil
}

//ChangeStatus changes task status
func (t *Task) ChangeStatus(status int) error {
	t.Status = status
	return nil
}
