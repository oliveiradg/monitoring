package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const loopDeMonitoramento = 5
const delay = 5

func main() {
	exibeIntrodução()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exinindo Logs...")
		case 0:
			fmt.Println("saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando")
			os.Exit(-1)
		}

	}

}

//-------------------- FUNÇÕES ---------------------------

func exibeIntrodução() {
	nome := "Joao"
	versao := 1.1

	fmt.Println("Olá Sr, ", nome)
	fmt.Println("Este Programa está na versão ", versao)

}

func exibeMenu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O Comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando....")
	sites := []string{"https://modalgr.com.br", "https://www.uol.com.br/",
		"https://globoplay.globo.com/", "https://random-status-code.herokuapp.com/"}

	fmt.Println(sites)

	for i := 0; i < loopDeMonitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site:", i, ": ", site)
			testaSite(site)

		}

		time.Sleep(delay * time.Second)

		fmt.Println("")

	}

	fmt.Println("")

}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O Site:", site, "Foi Carregado com Sucesso!")

	} else {
		fmt.Println("O Site:", site, "Está com Problemas. Status Code:", resp.StatusCode)
	}

}
