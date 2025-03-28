package main

import "fmt"

func media(n1, n2 int) bool {
	defer fmt.Println("Media calculada!!!Resultado será retornado")

	soma := (n1 + n2) / 2
	if soma >= 6 {
		return true

	} else {
		return false
	}
}

func main() {
	fmt.Println(media(6, 6))
}
