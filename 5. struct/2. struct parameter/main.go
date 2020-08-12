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

func main() {
	user1 := User{1, "Khaiz", "Tammam", "tamamizid@gmail.com", true}

	result := displayUser(user1)
	fmt.Println(result)
}

func displayUser(user User) string {
	result := fmt.Sprintf("Nama : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}
