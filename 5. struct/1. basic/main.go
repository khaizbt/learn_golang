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
	user := User{}
	//Mengisi Properties

	user.ID = 1
	user.FirstName = "Khaiz Badaru"
	user.LastName = "Tammam"
	user.Email = "Tamamizid@gmail.com"
	user.IsActive = true

	// Cara2
	user2 := User{ID: 2,
		FirstName: "Rodiyatun",
		LastName:  "Mardiyah",
		Email:     "Rodidiyatunmardiyah@gmail.com",
		IsActive:  true,
	}
	//Cara 3
	user3 := User{ID: 3, FirstName: "Muhammad", LastName: "Nathan", Email: "nathan@gmail.com", IsActive: false}

	//Cara 4
	user4 := User{4, "Khafif", "Damanhuri", "khafif@gmail.com", true}
	fmt.Println(user)
	fmt.Println(user2.FirstName)
	fmt.Println(user3)
	fmt.Println(user4)
}
