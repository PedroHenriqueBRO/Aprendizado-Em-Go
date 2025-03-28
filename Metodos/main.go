package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (us usuario) salvar() {
	fmt.Printf("Usuario salvo com sucesso : %s\n", us.nome)
}
func (us *usuario) addidade() {
	us.idade++
}

func main() {
	usuario1 := usuario{"Pedro", 30}
	usuario1.salvar()
	usuario1.addidade()
	fmt.Println(usuario1)

}
