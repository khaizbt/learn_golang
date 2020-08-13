package main

import "fmt"

func main() {
	number := 5
	fmt.Println("Nilai Awal adalah", number)
	fmt.Println("Alamat Memory pada number adalah", &number)

	change(&number, 100) //pengiriman ke function menggunakan alamat memory

	fmt.Println("Nilai akhir Adalah", number)
	fmt.Println("Alamat Nilai akhir adalah", &number)
}

func change(numberOld *int, numberNew int) { //karena memory jadinya depan type data memakai tanda *
	*numberOld = numberNew //perubahan nilai dari variabel yang mempunyai alamat memory yang sama

	fmt.Println("Alamat Nomor di Function adalah", *numberOld) //karena diatas masih dalam bentuk * maka isinya masih berupa string
	fmt.Println("Nomor di function adalah", numberOld)         //bentuknya sudah menjadi alamt memory karena kebalikan tanda *
	// fmt.
}
