package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3

func main() {
	exibeIntroducao()
	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	// fmt.Println("O tipo da variável versão é", reflect.TypeOf(versao))

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {

	// 	fmt.Println("Não conheço este comando")
	// }

}

func exibeIntroducao() {
	nome := "Ringo"
	var versao float32 = 1.20
	fmt.Println("Hello,", nome)
	fmt.Println("Versão do programa:", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 1; i <= monitoramentos; i += 1 {
		for i, site := range sites {
			fmt.Println("Testando site", i+1)
			testaSite(site)
		}
		time.Sleep(5 * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problemas")
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
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
	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
