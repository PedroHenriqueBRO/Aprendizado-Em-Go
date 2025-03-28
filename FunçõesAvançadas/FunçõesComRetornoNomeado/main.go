package main

import "fmt"

func somaesubtração(numb1 int, numb2 int) (soma int, subtracao int) {
	soma = numb1 + numb2
	subtracao = numb1 - numb2
	return //nao precisa colocar soma e subtração no return pois ele é nomeado
}

func main() {
	fmt.Println(somaesubtração(3, 2))

}
