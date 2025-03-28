package main

import (
	"errors"
	"fmt"
)

func main() {
	var bits8 int8 = 127
	var bits16 int16 = 1024*32 - 1
	var bits32 int32 = (1024 * 1024 * 1024 * 2) - 1
	var bits64 int64 = (1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 8) - 1
	fmt.Println(bits8, bits16, bits32, bits64)
	//unsigned int é um int sem sinal e serve para os mesmo tamanhos que os int mas só para a parte positiva

	//float
	var numerofracionario float32 = 120.34
	fmt.Println(numerofracionario)

	//string
	var nome string = "Pedro"
	fmt.Println(nome)

	//booleano
	var booleano bool = true
	fmt.Println(booleano)

	//erro
	var teste error = errors.New("Erro interno")
	fmt.Println(teste)

}
