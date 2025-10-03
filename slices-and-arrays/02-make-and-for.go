package slicesandarrays

import "fmt"

func MakeAndFor() {
	lista := []int{4, 9, 7}
	lista = append(lista, 19)

	listaDeString := []string{"Go Lang", "C#", "JavaScript"}

	listaDeString = append(listaDeString, "Solidity")

	fmt.Println("Lista de String: ", lista)

	fmt.Println("Lista: ", lista)
	fmt.Println("Lista: ", lista[0])
	fmt.Println("Lista: ", lista[1])
	fmt.Println("Lista: ", lista[2])
	fmt.Println("Tamanho da lista: ", len(lista))

}
