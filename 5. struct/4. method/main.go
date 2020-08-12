package main

import (
	"fmt"
)

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
	user := User{1, "Khaiz Badaru", "Tammam", "tamamizid@gmail.com", true}
	user2 := User{2, "Khafif", "Damanhuri", "khafif@gmail.com", true}

	result := user.display()

	fmt.Println(result)
	fmt.Println(user2.display())
}

func (user User) display() string { //Method dengan nama Display yang dimiliki oleh object user
	//Tanpa parameter, jadi otomatis di Parsing data dari main
	return fmt.Sprintf("Nama Saya adalah %s %s dengan email %s", user.FirstName, user.LastName, user.Email)
}
