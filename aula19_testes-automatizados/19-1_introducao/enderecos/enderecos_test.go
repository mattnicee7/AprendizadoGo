package enderecos

import "testing"

// teste de unidade
// Test+Nome da funcao
// arquivo precisa ter _test.go
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo recebido é diferente do esperado")
	}

}
