package main

import "fmt"

func main() {
	channeltarefa := make(chan int, 45)
	channelresultado := make(chan int, 45)
	go worker(channeltarefa, channelresultado)
	go worker(channeltarefa, channelresultado)
	go worker(channeltarefa, channelresultado)
	go worker(channeltarefa, channelresultado)

	for i := 0; i < 45; i++ {
		channeltarefa <- i
	}
	for i := 0; i < 45; i++ {
		resultado := <-channelresultado
		fmt.Println(i, resultado)
	}

}
func worker(enviar <-chan int, recebe chan<- int) {
	for i := 0; i < 45; i++ {
		recebe <- fibonacci(<-enviar)
	}
}
func fibonacci(posicao int) int {
	if posicao == 0 {
		return 0
	}
	if posicao <= 2 {
		return 1
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}
