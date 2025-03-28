package main

import (
	"fmt"
	"math"
)

// struct retangulo
type Retangulo struct {
	altura float64
	base   float64
}

// struct circulo
type Circulo struct {
	raio float64
}
type forma interface {
	area() float64
}

// funcao retangulo
func (r Retangulo) area() float64 {
	return r.altura * r.base
}
func retornaarea(tipo forma) {
	fmt.Println(tipo.area())

}

// funcao circulo
func (r Circulo) area() float64 {
	return (r.raio * r.raio) * math.Pi
}
func retornaqualquercoisa(tipo interface{}) {
	fmt.Println(tipo)
}

func main() {
	novo := Retangulo{4, 5}
	retornaarea(novo)
	retornaqualquercoisa(novo)
	fmt.Println(novo)

}
