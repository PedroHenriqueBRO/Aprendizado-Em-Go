package main

import "fmt"

func recuperarexecucao() {
	if n := recover(); n != nil {
		fmt.Println("Execução recuperada")
	}
}

func media(n1, n2 int) bool {
	defer recuperarexecucao()
	if (n1+n2)/2 > 6 {
		return true
	} else if (n1+n2)/2 < 6 {
		return false
	}
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(media(6, 6))
}
