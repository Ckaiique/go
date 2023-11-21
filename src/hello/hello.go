package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao() 
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "kaique"
	idade := 28
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Est programa está na vesão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {

	site := leSiteDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, sites := range site {
			fmt.Println("Testando Site:", i, ":", sites)
			testeSite(sites)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}
	fmt.Println("")

}

func testeSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso!")
		registraLogs(site , true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLogs(site , false)
	}
}

func leSiteDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		fmt.Println(linha)
		if err == io.EOF {
			break
		}

	}

	return sites
}

func registraLogs(site string, status bool){

	arquivo ,err := os.OpenFile("log.txt",os.O_RDWR|os.O_CREATE| os.O_APPEND, 0666)
	
	if err != nil {
		fmt.Println("Ocorreu um nos logs, erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - "+ site + "- Online: "+ strconv.FormatBool(status) +"\n" )
	arquivo.Close()
}


func imprimeLogs()  {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo) )

}