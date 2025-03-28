package main

import "fmt"

type usuario struct {
	name  string
	idade int8
}

func main() {
	var us1 usuario
	us1.name = "Pedro"
	us1.idade = 21
	fmt.Println(us1)
	us2 := usuario{"Joao", 32}
	fmt.Println(us2)

}
