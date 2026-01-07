package app

import "github.com/urfave/cli"

// gerar vai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de Linha de Comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	return app
}
