package enderecos

import "strings"

// TipoDeEndereco verifica se um endereco tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavraDoEndereco := strings.ToLower(strings.Split(endereco, " ")[0])

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			return strings.Title(tipo)
		}
	}

	return "Invalido"
}
