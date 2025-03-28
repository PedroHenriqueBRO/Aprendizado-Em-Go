package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("caramba@gmail.com")
	fmt.Println(erro)
}
