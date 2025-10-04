package main

import "fmt"

func main() {
	ImprimeMensagem("mensagem x")
	ImprimeMensagem("mensagem y")
}

func ImprimeMensagem(mensage string) {
	fmt.Println(mensage)
}
