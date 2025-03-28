package main

import (
	"fmt"
)

func main() {
	var (
		nome  string = "Pedro Henrique"
		idade int    = 21
	)
	var4, var5 := "oi", "18"
	fmt.Printf("Meu nome é %s e tenho %d anos de idade \n", nome, idade)
	fmt.Println(var4, var5)
	var4, var5 = var5, var4
	fmt.Println(var4, var5)

}
