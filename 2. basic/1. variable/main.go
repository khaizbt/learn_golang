package main

import (
	"fmt"
)

func main() {
	//Variable var definisi langsung
	var name string = "Golang"
	fmt.Println(name)

	//variable var tidak langsung
	var name01 string
	name01 = "Laravel"
	fmt.Println(name01)

	//variable integer
	var integer int = 100
	fmt.Println(integer)

	//variable boolean
	var boolean bool = true
	fmt.Println(boolean)

	//variable float
	var float float64 = 3.18
	fmt.Println(float)

	//variable dinamic
	nama := "Khaiz badaru Tammam"
	//mengubah variable nama
	nama = "Nama saya adalah tamam"
	fmt.Println(nama)

}
