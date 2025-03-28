package enderecos

import "testing"

type estruturadeteste struct {
	enderecoinserido  string
	resultadoesperado string
}

func TestAnalisaEndereco(t *testing.T) {

	teste := []estruturadeteste{{"Rua Francisco sa", "Rua"}, {"Avenida Doidos 2", "Avenida"}, {"Bairro Centro", "Bairro"}, {"", "Tipo inválido"}}

	for _, tipoteste := range teste {
		resultadorecebido := AnalisaEndereco(tipoteste.enderecoinserido)
		if resultadorecebido != tipoteste.resultadoesperado {
			t.Error("Resultado diferente do esperado")
		}
	}

}
