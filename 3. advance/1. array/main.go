package main

import (
	"fmt"
)

func main() {
	var language [5]string
	language[0] = "javascript"
	language[1] = "Go"
	language[2] = "PHP"
	language[3] = "Ruby"
	language[4] = "Python"

	//cara kedua
	nations := [5]string{
		"Indonesia", "Malaysia", "Japan", "Thailand", "Portugal",
	}

	//Jika tidak didefinisikan Elemen
	framework := [...]string{
		"Laravel",
		"Codeigniter",
		"Django",
		"Flask",
		"Springboot",
		"Yii",
	}

	//Foreach in array
	for _, word := range framework {
		fmt.Println(word)
	}

	fmt.Println("=====================")
	fmt.Println(nations)
	fmt.Println("=====================")
	fmt.Println(language)
	fmt.Println("=====================")
	length := len(framework)
	fmt.Println(length)
}
