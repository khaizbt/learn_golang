package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func main() {
	gamer := Gamer{Name: "Khaiz"}
	gamer.AddGame("PES 2020")
	gamer.AddGame("FIFA 2020")
	for _, game := range gamer.Games {
		fmt.Println(&game)
	}
}

func (gamer *Gamer) AddGame(game string) { //tanda Bintang agar dimasukan ke memory yang sama
	gamer.Games = append(gamer.Games, game)
}

