package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"bufio"
	"io"
	"strings"
	"strconv"
	"io/ioutil"
)


func main() {

	showIntro()

	for {
		showMenu()
		command := readCommand()
	
		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogger()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Giovani"
	version := 1.0
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	sites := readSitesData()

for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			siteTesting(site)
		}		
	}
	time.Sleep(time.Second * 10)
}

func siteTesting(site string) {
	response,err := http.Get(site) 
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		loggerRegister(site, true)
	} else {
		fmt.Println("Site:", site, "não foi carregado com sucesso")
		loggerRegister(site, false)
	}
}

func readSitesData() []string {
	var sites []string
	data, err := os.Open("sites.txt")


	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(data)
	
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	
	data.Close()
	return sites
}

func loggerRegister(site string, status bool) {
  data, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
			fmt.Println("Ocorreu um erro:", err)
	}

	data.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + 
	" - online: " + strconv.FormatBool(status) + "\n")

	data.Close()
}

func showLogger() {
	fmt.Println("Exibindo Logs...")
	data, err := ioutil.ReadFile("log.txt")

	if err != nil {
			fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(data))
}