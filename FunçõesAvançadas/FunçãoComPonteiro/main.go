package main

import "fmt"

func invertesinal(numb *int) {
	*numb = *numb * -1

}

func main() {
	numb := 5
	invertesinal(&numb)
	fmt.Println(numb)

}
