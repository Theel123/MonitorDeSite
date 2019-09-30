package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {

		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs")
		case 0:
			fmt.Println("Saindo.. BYE")
		default:
			fmt.Println("Digite uma opção valida")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "matheus"
	idade := 24
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade atualmente é ", idade, "anos")
	fmt.Println("Este programa este na versao", versao)

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	site := "https://www.alura.com.br"
	http.Get(site)
	res, _ := http.Get(site)
	fmt.Println(res)

	if res.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, "esta com problemas. Status Code: ", res.StatusCode)
	}
}

func devolveNomeIdade() (string, int) {
	nome := "matheus"
	idade := 19
	return nome, idade
}
