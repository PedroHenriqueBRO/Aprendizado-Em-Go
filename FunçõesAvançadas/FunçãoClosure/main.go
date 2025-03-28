package main

import "fmt"

func closure() func() {
	texto := "Texto2"

	novafunc := func() {
		fmt.Println(texto)
	}
	return novafunc
}

func main() {
	texto := "Texto1"
	fmt.Println(texto)

	novotexto := closure()
	novotexto()
}
