package main

import (
	"fmt"
)

func main() {
	gamingConsole := []string{"XBox", "PC", "Playstation"}

	fmt.Println(gamingConsole)

	//cara 2
	var anime []string

	anime = append(anime, "Kimi No Nawa")
	anime = append(anime, "Tenki no Ko")
	anime = append(anime, "Kimetsu No Yaiba")
	anime = append(anime, "Anohana")

	fmt.Println(anime)

	//mengambil slice dalam bentuk array
	fmt.Println(anime[0:2]) //Mengambil index 0-2

	//Mengambil semua slice dalam bentuk array
	fmt.Println(gamingConsole[:])
}
