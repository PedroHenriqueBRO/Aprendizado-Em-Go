package enderecos

import "strings"

//Função para analisar endereços
func AnalisaEndereco(endereco string) string {
	tipos := []string{"Rua", "Bairro", "Avenida"}
	nova := strings.Split(endereco, " ")[0]
	for _, tipo := range tipos {
		if tipo == nova {
			return tipo
		}
	}
	return "Tipo inválido"
}
