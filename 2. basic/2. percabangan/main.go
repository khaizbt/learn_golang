package main

import (
	"fmt"
)

func main() {
	score := 80
	var grade string

	if score == 80 {
		grade = "Adalah 80"
	} else {
		grade = "Bukan 80"
	}

	//elseif
	if score > 80 {
		grade = "Lebih dari 80"
	} else if score < 80 {
		grade = "Lebih dari 80"
	} else {
		grade = "NULL"
	}

	//if bercabang
	if score == 80 {
		if score > 70 {
			grade = "Maka itu 80"
		}
	}

	fmt.Println(grade)
}
