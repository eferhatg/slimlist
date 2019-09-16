package main

import (
	"fmt"

	"github.com/eferhatg/slimlist/slimlist"
)

func main() {

	task1()
	task2()
	task3()
	task4()

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

func task3() {
	slimlist.SetGlobalEmailSender()
	user1, _ := slimlist.NewUser("taskchenger@user.com")
	user2, _ := slimlist.NewUser("eferhatg@gmail.com")
	t, _ := user1.NewTask("Morning task")
	t.AssignToUser(*user2)
	user1.ChangeTaskStatusWithNotify(*slimlist.GlobalEmailSender, t, slimlist.InProgress)
	fmt.Println(t)
}

func task4() {

	admin, _ := slimlist.NewUser("admin@user.com")
	user, _ := slimlist.NewUser("regularuser@gmail.com")

	t, _ := admin.NewTask("Morning task")
	t.AssignToUser(*user)
	c, _ := user.AddComment("test_comment", t)
	err := admin.DeleteComment(c.ID, t)
	if err != nil {
		fmt.Println(err)
	}
	admin.SetAdmin()
	err = admin.DeleteComment(c.ID, t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
