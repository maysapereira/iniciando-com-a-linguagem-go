package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
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
                iniciaMonitoramento()
            case 2:
                fmt.Println("Exibindo logs...")
            case 0:
                fmt.Println("Saindo do programa...")
                os.Exit(0)
            default:
                fmt.Println("Escolha uma das opções válidas")
                os.Exit(-1)
    }

    }
}

func exibeIntroducao() {
    nome := "Maysa"
    versao := 1.3
    fmt.Println("Olá", nome, "!")

    fmt.Println("")

    fmt.Println("Este programa está na versão", versao)

    fmt.Println("")
}

func exibeMenu() {
    fmt.Println("1- Iniciar monitoramento")
    fmt.Println("2- Exibir logs")
    fmt.Println("0- Sair do programa")

    fmt.Println("")
}

func leComando() int {
    var comando int
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi o", comando)

    return comando
}

func iniciaMonitoramento() {
    fmt.Println("")

    fmt.Println("Iniciando monitoramento...")

    fmt.Println("")

    sites:= []string {"https://go.dev", "https://pagar.me", "https://www.stone.com.br", "https://github.com/maysapereira"}
    

for i := 0; i <monitoramentos; i++ {
    for i, site := range sites {
        fmt.Println("Testando site", i, ":", site)
        testaSites(site)
        fmt.Println("")
    }

    time.Sleep(delay * time.Minute)
}

    fmt.Println("")

}

func testaSites(site string) {
    resp, err := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
        registraLogs(site, true)
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:")
        registraLogs(site, false)
    }

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

}

func registraLogs(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)

    arquivo.WriteString(site + " - online: " + strconv.FormatBool(status) + "\n")

    arquivo.Close()
}