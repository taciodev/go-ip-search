package app

import "github.com/urfave/cli"

// Gerar aplicação
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor da internet"

	return app
}
