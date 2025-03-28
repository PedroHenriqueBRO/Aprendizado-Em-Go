package main

import "fmt"

func main() {
	func() {
		fmt.Println("Ola")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Ola mundo")

	resultado := func(texto string) string {
		return fmt.Sprintf("Resultado : %s", texto)
	}("Pedro")
	fmt.Println(resultado)
}
