package main

import "fmt"

func identificaChá(inteiro int) {
	var cont int8 = 0
	var participantes [5]int
	for i := 0; i < 5; i++ {
		var aux int
		fmt.Scan(&aux)
		participantes[i] = aux
	}

	for _, suposi := range participantes {
		if suposi == inteiro {
			cont++
		}
	}
	fmt.Println(cont)
}

func main() {
	var inteiro int
	fmt.Scan(&inteiro)
	identificaChá(inteiro)
}
