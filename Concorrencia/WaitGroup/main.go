package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2) //adiciono duas rotinas a serem esperadas
	go func() {
		escrever("Pedro")
		waitgroup.Done() //faço e concluo uma
	}()
	go func() {
		escrever("Henrique")
		waitgroup.Done() //faço e concluo uma
	}()
	waitgroup.Wait() //espero as duas serem concluídas para terminar o programa

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
