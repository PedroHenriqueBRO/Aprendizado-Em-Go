package formas

import (
	"math"
	"testing"
)

type testecirculos struct {
	circulo      Circulo
	areaesperada float64
}
type testeretangulos struct {
	retangulo    Retangulo
	areaesperada float64
}

func TestAreas(t *testing.T) {
	t.Run("Circulo", func(t *testing.T) {
		testes := []testecirculos{{Circulo{Raio: 4.0}, 16 * math.Pi}, {Circulo{Raio: 5.0}, 25 * math.Pi}, {Circulo{Raio: 6.0}, 36 * math.Pi}}
		for _, teste := range testes {
			resultado := teste.circulo.Area()
			if resultado != teste.areaesperada {
				t.Fatalf("Área diferente da esperada")
			}
		}

	})
	t.Run("Retangulo", func(t *testing.T) {
		testes := []testeretangulos{{Retangulo{4.0, 5.0}, 20.0}, {Retangulo{4.0, 6.0}, 24.0}, {Retangulo{10.0, 5.0}, 50.0}}
		for _, teste := range testes {
			resultado := teste.retangulo.Area()
			if resultado != teste.areaesperada {
				t.Fatalf("Área diferente da esperada")
			}
		}

	})
}
