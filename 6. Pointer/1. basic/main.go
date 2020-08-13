package main

import (
	"fmt"
)

func main() {
	pointerA := 5
	pointerB := &pointerA //PointerB merujuk pada alamat memori dimana pointerA disimpan
	//dan tanda & adalah tanda untuk mengambil referencing/alamat memory dimana pointer A disimpan(referencing)

	fmt.Println(pointerA)
	fmt.Println("Ini adalah Alamat memory Pointer A :", pointerB, ", Yang didefinisikan di PointerB")

	fmt.Println("Menggunakan Tanda '*' untuk mengubah dari pointer ke int", *pointerB)
	//Untuk mengubah alamat memory ke nilai maka bisa menggunakan tanda * (deferencing)

	*pointerB = 10
	fmt.Println("Untuk mengubah nilainya juga memakai tanda * pada variablenya", *pointerB)

	fmt.Println("karena disimpan di Memory yang sama maka Pointer A juga berubah menjadi", pointerA)

	fmt.Println("################### CARA 2 #############################")

	var stringA string = "Aku Sayang Kamu"
	var stringB *string = &stringA //tanda bintang di type data digunakan untuk menyimpan/menampung alamat memory

	fmt.Println(stringA)
	fmt.Println(stringB)

	stringA = "Aku Benci kamu"

	fmt.Println(stringA)
	fmt.Println(stringB) //Alamat memory akan tetap sama walaupun valuenya diubah

}
