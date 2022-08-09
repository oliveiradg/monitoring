package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

	sites := leSitesDoArquivo()

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

	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O Site:", site, "Foi Carregado com Sucesso!")
		registraLog(site, true)

	} else {
		fmt.Println("O Site:", site, "Está com Problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + fmt.Sprint(status) + "\n")
	arquivo.Close()
}
