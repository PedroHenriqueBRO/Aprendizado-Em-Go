package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Pedro")
	escrever("Henrique")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
