package slimlist

//Keeps task status
const (
	New = iota
	InProgress
	Completed
	Archived
)

var tasks []Task

//Task keeps task definition
type Task struct {
	ID          string
	Description string
	Status      int
	Users       []User
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
