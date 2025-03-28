package main

import "fmt"

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}
func loucura(i1, i2 int32) (string, int32) {
	return "idade", i1 + i2
}

func main() {
	var numero1 int = 32
	var numero2 int = 64
	fmt.Println(somar(numero1, numero2))
	soma := somar(numero1, numero2)
	fmt.Println(soma)

	var j = func(txt string) {
		fmt.Println(txt)

	}
	j("pedro")
	fmt.Println(loucura(20, 1))

}
