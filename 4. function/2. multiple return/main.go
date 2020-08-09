package main

import (
	"fmt"
)

func main() {
	luas, keliling := calculate(10, 12) //jika misal keliling tidak digunakan bisa memakai tanda _

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println("====================")

	kali, bagi := perkalianPembagian(45, 5)

	fmt.Println(kali)
	fmt.Println(bagi)
}

func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

//Predifined Return Valuee
func perkalianPembagian(number, numberTwo int) (kali int, bagi int) { //variabel didefinisikan di Output
	kali = number * numberTwo
	bagi = number / numberTwo

	return //cukup dengan return saja, karena otomatis kereturn semua
}
