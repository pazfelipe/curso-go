package main

import "fmt"

func executaCallback(f func(str string), str string) {
	f(str)
}

func meuCallback(str string) {
	fmt.Println(str)
}

func executaCallbackSemParametros(f func()) {
	f()
}

func meuCallbackSemParamentro() {
	fmt.Println("Sem argumento")
}

func main() {

	executaCallback(meuCallback, "Executando um callback")
	executaCallbackSemParametros(meuCallbackSemParamentro)

}
