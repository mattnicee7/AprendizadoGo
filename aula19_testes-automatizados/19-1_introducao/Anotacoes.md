Executar todos os testes do pacote ou subpastas:
go test ./...

Executar testes em modo verboso:
go test -v

Executar testes ignorando cache:
go test -count=1

Executar testes com cobertura:
go test -cover

Gerar arquivo de cobertura:
go test -coverprofile=cobertura.txt

Exibir cobertura por funcao no terminal:
go tool cover -func=cobertura.txt

Visualizar cobertura em HTML no navegador:
go tool cover -html=cobertura.txt

Observacao: coverage em Go mede statements executados, nao todos os cenarios de negocio.