package main

import (
	"fmt"
)

func main() {
	number := 5

	switch number {
	case 1: //Jika Number sama dengan 1
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("Tiga")
	default: //sama dengan Else pada if else
		fmt.Println("Tidak termasuk angka manapun")
	}
}
