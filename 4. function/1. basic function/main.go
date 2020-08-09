package main

import "fmt"

func main() {
	printMyResult()
	printParameter("saya print parameter")

	sentence := printOutput("saya sedang")
	fmt.Println(sentence)
	fmt.Println(add(17, 10))
}

func printMyResult() {
	fmt.Println("Saya belajar golang")
}

//function dengan parameter
func printParameter(parameter string) {
	fmt.Println(parameter)
}

//function dengan output
func printOutput(parameter string) string {
	sentence := parameter + " Belajar Golang"
	return sentence
}

//fungsi aritmatika
func add(number, numberTwo int) int { //Jika parameter type datanya sama maka bisa diinisialisasikan 1 type data
	return number + numberTwo
}
