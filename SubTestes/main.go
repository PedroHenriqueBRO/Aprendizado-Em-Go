package main

import (
	"fmt"
	"testes-formas/formas"
)

func main() {
	novo := formas.Circulo{Raio: 4}
	fmt.Println(novo.Area())

}
