package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Ola"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Mundo"
		}
	}()

	for {
		select {
		case mensagem1 := <-channel1:
			fmt.Println(mensagem1)
		case mensagem2 := <-channel2:
			fmt.Println(mensagem2)
		}
	}
}
