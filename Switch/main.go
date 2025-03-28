package main

import "fmt"

func main() {
	var numb int = 5
	switch {
	case numb < 6:
		fmt.Println("5")
		numb--
		fallthrough
	case numb < 5:
		fmt.Println("4")
		numb--
		fallthrough
	case numb < 4:
		fmt.Println("3")
		numb--
		fallthrough
	case numb < 3:
		fmt.Println("2")
		numb--
		fallthrough
	case numb < 2:
		fmt.Println("1")
		numb--
		fallthrough
	case numb < 1:
		fmt.Println("0")
	default:
		break

	}
}
