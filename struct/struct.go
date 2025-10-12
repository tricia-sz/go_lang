package struc

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
}

func Struuct() {
	fmt.Println("Iniciando...")

	endereco := endereco{
		rua:    "Olavo Nundes",
		numero: 13,
		cidade: "Campo Grande",
	}

	fmt.Println(endereco)

	endereco.numero = 18
	fmt.Println(endereco.numero)
}
