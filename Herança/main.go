package main

import "fmt"

type pessoa struct {
	name   string
	idade  int8
	altura float32
}
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	var es1 estudante = estudante{pessoa{"Pedro", 21, 1.78}, "CC", "IFNMG"}
	fmt.Println(es1)
}
