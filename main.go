package main

import (
	"fmt"
	model "learngo/model"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Olavo Nundes",
		Numero: 13,
		Cidade: "Campo Grande",
	}

	pessoa := model.Pessoa{
		Nome:     "Tricia",
		Endereco: endereco,
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)

	endereco.Numero = 18
	fmt.Println(endereco.Numero)
}
