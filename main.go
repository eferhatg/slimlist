package main

import (
	"fmt"

	"github.com/eferhatg/slimlist/slimlist"
)

/*
Simulating the cases one by one
*/
func main() {

	task1()
	task2()
	task3()
	task4()

}

/*
Develop a task management system. Each task should have its own status,
description, and users assigned to it.
Statuses:  new, in progress, completed, archived.
*/
func task1() {
	user1, _ := slimlist.NewUser("task1user1@user.com")
	user2, _ := slimlist.NewUser("task1user2@user.com")
	t, _ := user1.NewTask("Morning task")
	t.AssignToUser(*user2)
	fmt.Println(t)
}

/*
The system should provide the user roles. Only the administrator should be able to archive tasks.
*/
func task2() {
	user1, _ := slimlist.NewUser("task2user1@user.com")
	user2, _ := slimlist.NewUser("task2user2@user.com")
	t, _ := user1.NewTask("Morning task")
	t.AssignToUser(*user2)
	user1.SetAdmin()
	user1.ChangeTaskStatus(t, slimlist.Archived)
	fmt.Println(t)
}

/*
Users assigned to tasks receive notifications about changes by email.
*/
func task3() {
	slimlist.SetGlobalEmailSender()
	user1, _ := slimlist.NewUser("taskchenger@user.com")
	user2, _ := slimlist.NewUser("eferhatg@gmail.com")
	t, _ := user1.NewTask("Morning task")
	t.AssignToUser(*user2)
	user1.ChangeTaskStatusWithNotify(*slimlist.GlobalEmailSender, t, slimlist.InProgress)
	fmt.Println(t)
}

/*
Users should be able to comment on the tasks. Only administrators can delete comments.
*/
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
