package main

import (
	"fmt"

	"github.com/eferhatg/slimlist/slimlist"
)

func main() {
	task1()

}

func task1() {
	user1, _ := slimlist.NewUser("user1@user.com")
	user2, _ := slimlist.NewUser("user2@user.com")
	t, _ := user1.NewTask("Morning task")
	t.AssignToUser(*user2)
	fmt.Println(t)
}
