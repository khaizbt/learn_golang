package main

import "fmt"

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

//Mendisplay Group from Embedded Struct
func main() {
	user := User{1, "Khafif", "Damanhuri", "khafif@gmail.com", false}
	user1 := User{1, "Khaiz", "Tammam", "tamamizid@gmail.com", true}

	users := []User{user, user1}

	group := Group{"Tester", user, users, true}

	group.displayGroup()
}

func (group Group) displayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Nama User :")

	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
