package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Ola"
	channel <- "Ola"
	channel <- "Ola" //da deadlock pois o buffer é de tamanho 2

	mensagem := <-channel

	fmt.Println(mensagem)
}
