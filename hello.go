package main

import "fmt"
import "os"

func main() {

    exibeIntroducao()

    exibeMenu()

    comando := leComando()

    switch comando {
        case 1:
            fmt.Println("Iniciando monitoramento...")
        case 2:
            fmt.Println("Exibindo logs...")
        case 0:
            fmt.Println("Saindo do programa...")
            os.Exit(0)
        default:
            fmt.Println("Escolha uma das opções válidas")
            os.Exit(-1) //caso dê erro ao sair do programa
    }
}

func exibeIntroducao() {
    nome := "Maysa"
    versao := 1.1
    fmt.Println("Olá, sra.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
    fmt.Println("1- Iniciar monitoramento")
    fmt.Println("2- Exibir logs")
    fmt.Println("0- Sair do programa")
}

func leComando() int {
    var comando int
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi o", comando)

    return comando
}