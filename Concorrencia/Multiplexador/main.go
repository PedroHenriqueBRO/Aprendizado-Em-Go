package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Pedro"), escrever("Henrique"))
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func multiplexar(channel1 <-chan string, channel2 <-chan string) <-chan string {
	canalsaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-channel1:
				canalsaida <- mensagem
			case mensagem := <-channel2:
				canalsaida <- mensagem

			}
		}
	}()
	return canalsaida

}

func escrever(texto string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- texto
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel

}
