package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("belajar golang hari ke ", i)
	}

	//Karena di Go lang tidak ada while maka bisa dengan cara ini
	y := 1
	for y <= 100 {
		fmt.Println("Bekerja backend hari ke ", y)
		y++
	}

	//for untuk mengambil dari semua collection
	title := "Golang adalah best programming language"

	for index, letter := range title { //Index adalah index ke - n dan letter adalah pemecahan kata-katanya
		fmt.Println("Index ", index, "Letter ", string(letter))
	}
}
