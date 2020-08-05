package main

import "fmt"

func main() {
	var myMap map[string]int //deklarasi variable,  default value  dari map adalah nil

	myMap = map[string]int{} //namun karena deafult value darin map adalah nil maka perlu inisialisasi nilai default denganditambahkan {} diakhir

	myMap["ruby"] = 9
	myMap["Javascript"] = 5

	fmt.Println(myMap)
	fmt.Println(myMap["ruby"])
	fmt.Println(myMap["php"]) //default value adalah 0 karena belum ada item yang tersimpan dalam php

	//atau bisa melihat value dengan cara ini
	value, isAvailable := myMap["python"]
	fmt.Println(isAvailable, value) //isAvailable akan keluar boolean, itu karena default dari Golangnya

	fmt.Println("================")
	//Cara 2
	lang := map[string]string{
		"Indonesia": "Is Beatutiful",
		"Java":      "Is our programming language", //Aturan vertikal sama seperti array
		"English":   "Is Eastern Language",
	} //Pemisah antara Index dengan Value adalah tanda titik dua (:)

	for key, value := range lang {
		fmt.Println("Key ", key, " value ", value)
	}

	fmt.Println("================")

	//Hapus map
	delete(lang, "Java")

	for key, value := range lang {
		fmt.Println("Key ", key, " value ", value)
	}
}
