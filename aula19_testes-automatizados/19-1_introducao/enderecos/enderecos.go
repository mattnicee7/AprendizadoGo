package enderecos

import "strings"

// TipoDeEndereco verifica se um endereco tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	endereco = strings.TrimSpace(strings.ToLower(endereco))
	if endereco == "" {
		return "Invalido"
	}

	primeiraPalavra := strings.Fields(endereco)[0]

	switch primeiraPalavra {
	case "rua":
		return "Rua"
	case "avenida":
		return "Avenida"
	case "estrada":
		return "Estrada"
	case "rodovia":
		return "Rodovia"
	default:
		return "Invalido"
	}
}
