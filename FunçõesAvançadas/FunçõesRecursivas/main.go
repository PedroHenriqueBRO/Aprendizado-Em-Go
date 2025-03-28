package main

import "fmt"

var M [1001]int
var J [1001]bool

func fibonacci(posicao int) int {
	if posicao == 0 {
		return 0
	}
	if posicao <= 2 {
		return 1
	}
	if J[posicao] {
		return M[posicao]
	}
	M[posicao] = fibonacci(posicao-1) + fibonacci(posicao-2)
	J[posicao] = true
	return M[posicao]
}

func main() {
	fmt.Println(fibonacci(1000))
}
