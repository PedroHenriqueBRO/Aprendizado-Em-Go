package main

import (
	"fmt"
)

func main() {
	var matriz [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matriz[i][j] = j
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}

}
