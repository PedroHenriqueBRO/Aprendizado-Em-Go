package main

import "fmt"

func fibonnaci(inteiro int, chamadas *int) int {
	if inteiro == 0 {
		return 0
	}
	if inteiro == 1 {
		return 1
	}
	*chamadas += 2
	return fibonnaci(inteiro-1, chamadas) + fibonnaci(inteiro-2, chamadas)

}

func main() {
	var chamadas int = 0
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var aux int
		fmt.Scan(&aux)
		resultado := fibonnaci(aux, &chamadas)
		fmt.Printf("fib(%d) = %d calls = %d\n", aux, chamadas, resultado)
		chamadas = 0

	}

}
