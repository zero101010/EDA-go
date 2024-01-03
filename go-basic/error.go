package main

import (
	"fmt"
	"log"
	"os"
)

// Podemos usar essas formas de lidar com erro:
// fmt.Println()
// log.Println() ---> Tem o timestramp
// log.SetOutput() --> Escreve o erro em um Writter, seja stdout ou em um arquivo
// log.Fatalln() ---> Retorna um Exit(1) e não roda os defer()
// panic() --> A execução é parada imediatamente, mas os defers ainda são executados
// log.Panicln() --> Tem a mesma função do panic, mas retorna o log, é o mais recomendado

func showFmtPrintln() {
	fmt.Println("fmt.Println: O pacote mais simples para erro")

}
func showLogPrintln() {
	log.Println("log.Println: Imprime o erro com o timestamp")

}

func showLogSetOutput() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// Salva em algum writter, poderia ser o Stdout ou o arquivo
	log.SetOutput(f)

	f2, err := os.Open("no-file")
	if err != nil {
		log.Println(err)
	}
	defer f2.Close()
}

func showPanic() {
	f, err := os.Open("no-file")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
func showLogPanicln() {
	f, err := os.Open("no-file")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()
}
func main() {
	showFmtPrintln()
	showLogPrintln()
	showLogSetOutput()
	// showPanic()
	showLogPanicln()
}
