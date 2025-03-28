package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go func() {
		escrever("Pedro", channel)
	}()
	/*for {

		mensagem, aberto := <-channel
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/
	for mensagem := range channel {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto
		time.Sleep(time.Second)
	}
	close(channel)
}
