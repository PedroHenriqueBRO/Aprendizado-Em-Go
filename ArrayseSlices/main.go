package main

import (
	"fmt"
)

func main() {

	var numbs [5]int8
	numbs[0] = 1
	numbs[1] = 2
	numbs[2] = 3
	numbs[3] = 4
	numbs[4] = 5
	for _, numb := range numbs {
		fmt.Println(numb)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(numbs[i])
	}
	fishs := [...]string{"aaa", "bbb", "ccc"}           //array sem tamanho definido na declaração mas que tem tamanho fixo de acordo a quantidade de itens nas chaves
	sharks := []string{"Hammerhead", "White", "Duende"} //slice
	for _, shark := range sharks {
		fmt.Println(shark)
	}
	sharks = append(sharks, "GG")
	sharks = sharks[1:3]
	fmt.Println(sharks)
	fmt.Println(fishs)

}
