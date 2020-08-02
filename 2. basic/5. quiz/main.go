package main

import (
	"fmt"
)

func main() {
	title := "Golang adalah best programming language"

	for index, letter := range title { //Index adalah index ke - n dan letter adalah pemecahan kata-katanya
		if index%2 == 0 {
			fmt.Println("Letter ", string(letter))
		}
	}
	fmt.Println("===========================")
	for _, letter := range title { //Underscore adalah syntax deklarasi jika dibawahnya tidak dipakai

		switch string(letter) {
		case "a", "i", "u", "e", "o":
			fmt.Println("Letter ", string(letter))

		}
	}
}
