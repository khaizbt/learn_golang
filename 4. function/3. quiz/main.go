package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{100, 80, 70, 54, 90}

	fmt.Println(sumArray(scores))
	result, err := calculate(10, 12, "?")

	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

func sumArray(score []int) int {

	var sum int
	for _, total := range score {
		sum = sum + total
	}

	return sum
}

func calculate(numberOne int, numberTwo int, operator string) (total int, errorResult error) {

	switch operator {
	case "+":
		total = numberOne + numberTwo
	case "-":
		total = numberOne - numberTwo
	case "*":
		total = numberOne * numberTwo
	case "/":
		total = numberOne / numberTwo
	default:
		errorResult = errors.New("Unknown Operation")
	}

	return
}
