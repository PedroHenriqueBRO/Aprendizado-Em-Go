package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("texto")

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
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
