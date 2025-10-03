// 21. Index
package slicesandarrays

import "fmt"

func Index() {
	listaToda := []int{2, 10, 9, 4, 8, 1, 3}
	segundaLista := listaToda[:3]
	terceiraLista := listaToda[4:]
	ultimoItem := listaToda[len(listaToda)-1:]

	fmt.Println(segundaLista)
	fmt.Println(terceiraLista)
	fmt.Println(ultimoItem)
}
