package main

import "fmt"

func variatica(numeros ...int) {
	fmt.Println(numeros) //numeros é um slice
	numeros = numeros[1:3]

	for _, numb := range numeros {
		fmt.Println(numb)
	}

}

func main() {
	variatica(1, 2, 3, 4, 5, 6)
}
