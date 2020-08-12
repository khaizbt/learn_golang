package main

import (
	"fmt"
)

//Embedded Struct adalah struct untuk struct lain
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	user := User{1, "Khafif", "Damanhuri", "khafif@gmail.com", false}
	user1 := User{1, "Khaiz", "Tammam", "tamamizid@gmail.com", true}

	users := []User{user, user1}

	group := Group{"Tester", user, users, true}

	displayGroup(group)
}

func displayUser(user User) string {
	result := fmt.Sprintf("Nama : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Printf("")
	fmt.Println("Nama User :")

	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
