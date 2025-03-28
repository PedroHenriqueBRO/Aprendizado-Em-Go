package main

import "fmt"

func appendarray(vetor []int, newitem int) []int {
	if cap(vetor) == len(vetor) {
		vetornovo := make([]int, len(vetor)+1, cap(vetor)*2)
		vetornovo = copia(vetornovo, vetor)
		vetornovo[len(vetor)] = newitem
		return vetornovo
	} else {
		vetornovo := make([]int, len(vetor)+1, cap(vetor))
		vetornovo = copia(vetornovo, vetor)
		vetornovo[len(vetor)] = newitem
		return vetornovo
	}

}
func copia(vetor []int, vetor2 []int) []int {

	for i := 0; i < len(vetor2); i++ {
		vetor[i] = vetor2[i]
	}
	return vetor

}

func main() {
	sharks := []int{1, 2, 3, 4, 5}
	fmt.Println(sharks)
	fmt.Println(len(sharks))
	fmt.Println(cap(sharks))
	sharks = appendarray(sharks, 6)
	fmt.Println(sharks)
	fmt.Println(len(sharks))
	fmt.Println(cap(sharks))
	sharks = appendarray(sharks, 7)
	fmt.Println(sharks)
	fmt.Println(len(sharks))
	fmt.Println(cap(sharks))
	sharks = appendarray(sharks, 8)
	fmt.Println(sharks)
	fmt.Println(len(sharks))
	fmt.Println(cap(sharks))
	sharks = appendarray(sharks, 9)
	fmt.Println(sharks)
	fmt.Println(len(sharks))
	fmt.Println(cap(sharks))
	sharks = appendarray(sharks, 10)
	fmt.Println(sharks)
	fmt.Println(len(sharks))
	fmt.Println(cap(sharks))
	sharks = appendarray(sharks, 11)
	fmt.Println(sharks)
	fmt.Println(len(sharks))
	fmt.Println(cap(sharks))

}
