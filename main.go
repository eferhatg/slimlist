package main

import (
	"fmt"

	"github.com/eferhatg/slimlist/slimlist"
)

func main() {
	task1()
	task2()

}

func task1() {
	user1, _ := slimlist.NewUser("task1user1@user.com")
	user2, _ := slimlist.NewUser("task1user2@user.com")
	t, _ := user1.NewTask("Morning task")
	t.AssignToUser(*user2)
	fmt.Println(t)
}

func task2() {
	user1, _ := slimlist.NewUser("task2user1@user.com")
	user2, _ := slimlist.NewUser("task2user2@user.com")
	t, _ := user1.NewTask("Morning task")
	t.AssignToUser(*user2)
	user1.SetAdmin()
	user1.ChangeTaskStatus(t, slimlist.Archived)
	fmt.Println(t)
}
