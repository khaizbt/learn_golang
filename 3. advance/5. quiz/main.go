package main

import (
	"fmt"
)

func main() {
	scores := []int{100, 80, 70, 54, 90}

	var total int

	for _, score := range scores {
		total = total + score
	}

	fmt.Println(float64(total) / float64(len(scores)))

	//mengambil goodscore(lebih atau sama dengan dari 80)

	var goodScore []int

	for _, score := range scores {
		if score >= 80 {
			goodScore = append(goodScore, score)
		}
	}

	fmt.Println(goodScore)
}
