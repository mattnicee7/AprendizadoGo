package enderecos_test

import (
	. "AprendizadoGo/aula19_testes-automatizados/19-1_introducao/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// teste de unidade
// Test+Nome da funcao
// arquivo precisa ter _test.go

// Pra executar os testes em todos arquivos do pacote:
// go test ./...

// Pra executar com modo verboso
// go test -v

// Pra verificar se ta checando todas possibilidades
// go test --cover
// go tool cover --func=cobertura.txt
// pra ver em html:
// go tool cover --html=cobertura.txt
func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		//{"Rua Abc", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Invalido"},
		{"Estrada qualquer", "Estrada"},
		//{"rua dos bobos", "Rua"},
		{" ", "Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
