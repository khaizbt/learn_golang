package main

import "fmt"

func main() {
	students := []map[string]string{ //slice yang isinya map yang ,mana keynya string valuenya string
		{"name": "Khaiz Badaru Tammam", "score": "A"}, //Map
		{"name": "Khafif Damanhuri", "score": "B"},
		{"name": "Banu", "score": "C"},
	}
	for _, student := range students {
		fmt.Println(student["name"], student["score"])
	}

}
