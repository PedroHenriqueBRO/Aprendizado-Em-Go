package formas

import "math"

//struct retangulo
type Retangulo struct {
	Altura float64
	Base   float64
}

//struct circulo
type Circulo struct {
	Raio float64
}

//funcao retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Base
}

//funcao circulo
func (r Circulo) Area() float64 {
	return (r.Raio * r.Raio) * math.Pi
}
